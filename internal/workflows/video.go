package workflows

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zibbp/ganymede/ent"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/queue"
	"github.com/zibbp/ganymede/internal/activities"
	"github.com/zibbp/ganymede/internal/database"
	"github.com/zibbp/ganymede/internal/dto"
	"github.com/zibbp/ganymede/internal/notification"
	"github.com/zibbp/ganymede/internal/utils"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func checkIfTasksAreDone(input dto.ArchiveVideoInput) error {
	log.Debug().Msgf("checking if tasks are done for video %s", input.VideoID)
	q, err := database.DB().Client.Queue.Query().Where(queue.ID(input.Queue.ID)).Only(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("error getting queue item")
		return err
	}

	if input.Queue.LiveArchive {
		if q.TaskVideoDownload == utils.Success && q.TaskVideoConvert == utils.Success && q.TaskVideoMove == utils.Success && q.TaskChatDownload == utils.Success && q.TaskChatConvert == utils.Success && q.TaskChatRender == utils.Success && q.TaskChatMove == utils.Success {
			log.Debug().Msgf("all tasks for video %s are done", input.VideoID)

			_, err := q.Update().SetVideoProcessing(false).SetChatProcessing(false).SetProcessing(false).Save(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("error updating queue item")
				return err
			}

			_, err = database.DB().Client.Vod.UpdateOneID(input.Vod.ID).SetProcessing(false).Save(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("error updating vod")
				return err
			}

			notification.SendLiveArchiveSuccessNotification(input.Channel, input.Vod, input.Queue)
		}
	} else {
		if q.TaskVideoDownload == utils.Success && q.TaskVideoConvert == utils.Success && q.TaskVideoMove == utils.Success && q.TaskChatDownload == utils.Success && q.TaskChatRender == utils.Success && q.TaskChatMove == utils.Success {
			log.Debug().Msgf("all tasks for video %s are done", input.VideoID)

			_, err := q.Update().SetVideoProcessing(false).SetChatProcessing(false).SetProcessing(false).Save(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("error updating queue item")
				return err
			}

			_, err = database.DB().Client.Vod.UpdateOneID(input.Vod.ID).SetProcessing(false).Save(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("error updating vod")
				return err
			}

			notification.SendVideoArchiveSuccessNotification(input.Channel, input.Vod, input.Queue)
		}
	}

	return nil
}

func workflowErrorHandler(err error, input dto.ArchiveVideoInput, task string) error {
	notification.SendErrorNotification(input.Channel, input.Vod, input.Queue, task)

	return err
}

// *Top Level Workflow*
func ArchiveVideoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{})

	// create directory
	err := workflow.ExecuteChildWorkflow(ctx, CreateDirectoryWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	// download thumbnails
	err = workflow.ExecuteChildWorkflow(ctx, DownloadTwitchThumbnailsWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	// save video info
	err = workflow.ExecuteChildWorkflow(ctx, SaveTwitchVideoInfoWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	// archive video
	videoFuture := workflow.ExecuteChildWorkflow(ctx, ArchiveTwitchVideoWorkflow, input)

	if input.Queue.ChatProcessing {
		chatFuture := workflow.ExecuteChildWorkflow(ctx, ArchiveTwitchChatWorkflow, input)
		if err := chatFuture.Get(ctx, nil); err != nil {
			return err
		}
	}

	if err := videoFuture.Get(ctx, nil); err != nil {
		return err
	}

	return nil
}

// *Top Level Workflow*
func ArchiveLiveVideoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{})

	// create directory
	err := workflow.ExecuteChildWorkflow(ctx, CreateDirectoryWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	// download thumbnails
	err = workflow.ExecuteChildWorkflow(ctx, DownloadTwitchLiveThumbnailsWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	// save video info
	err = workflow.ExecuteChildWorkflow(ctx, SaveTwitchLiveVideoInfoWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	var chatFuture workflow.ChildWorkflowFuture
	if input.Queue.ChatProcessing {
		chatFuture = workflow.ExecuteChildWorkflow(ctx, ArchiveTwitchLiveChatWorkflow, input)
		var chatWorkflowExecution workflow.Execution
		_ = chatFuture.GetChildWorkflowExecution().Get(ctx, &chatWorkflowExecution)

		log.Debug().Msgf("Live chat archive workflow ID: %s", chatWorkflowExecution.ID)
		input.LiveChatArchiveWorkflowId = chatWorkflowExecution.ID

		// execute chat download first to get a workflow ID for signals
		// the actual download of chat is held until the video is about to start
		liveChatFuture := workflow.ExecuteChildWorkflow(ctx, DownloadTwitchLiveChatWorkflow, input)
		var liveChatWorkflowExecution workflow.Execution
		_ = liveChatFuture.GetChildWorkflowExecution().Get(ctx, &liveChatWorkflowExecution)

		log.Debug().Msgf("Live chat workflow ID: %s", liveChatWorkflowExecution.ID)
		input.LiveChatWorkflowId = liveChatWorkflowExecution.ID
	}

	// archive video
	videoFuture := workflow.ExecuteChildWorkflow(ctx, ArchiveTwitchLiveVideoWorkflow, input)

	if err := videoFuture.Get(ctx, nil); err != nil {
		return err
	}

	if input.Queue.ChatProcessing {
		if err := chatFuture.Get(ctx, nil); err != nil {
			return err
		}
	}

	return nil
}

// *Low Level Workflow*
func CreateDirectoryWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	err := workflow.ExecuteActivity(ctx, activities.CreateDirectory, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "create-directory")
	}

	return nil
}

// *Low Level Workflow*
func DownloadTwitchThumbnailsWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	err := workflow.ExecuteActivity(ctx, activities.DownloadTwitchThumbnails, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "download-thumbnails")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func DownloadTwitchLiveThumbnailsWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	err := workflow.ExecuteActivity(ctx, activities.DownloadTwitchLiveThumbnails, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "download-thumbnails")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func SaveTwitchVideoInfoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	err := workflow.ExecuteActivity(ctx, activities.SaveTwitchVideoInfo, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "save-video-info")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func SaveTwitchLiveVideoInfoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	err := workflow.ExecuteActivity(ctx, activities.SaveTwitchLiveVideoInfo, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "save-video-info")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Mid Level Workflow*
func ArchiveTwitchVideoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {

	err := workflow.ExecuteChildWorkflow(ctx, DownloadTwitchVideoWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteChildWorkflow(ctx, PostprocessVideoWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteChildWorkflow(ctx, MoveVideoWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

// *Mid Level Workflow*
func ArchiveTwitchLiveVideoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {

	err := workflow.ExecuteChildWorkflow(ctx, DownloadTwitchLiveVideoWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteChildWorkflow(ctx, PostprocessVideoWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteChildWorkflow(ctx, MoveVideoWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil

}

// *Mid Level Workflow*
func ArchiveTwitchLiveChatWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	// download happened earlier, this is post-download tasks

	var signal utils.ArchiveTwitchLiveChatStartSignal
	signalChan := workflow.GetSignalChannel(ctx, "continue-chat-arhive")
	signalChan.Receive(ctx, &signal)

	log.Info().Msgf("Received signal: %v", signal)

	err := workflow.ExecuteChildWorkflow(ctx, ConvertTwitchLiveChatWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	if input.Queue.RenderChat {
		err = workflow.ExecuteChildWorkflow(ctx, RenderTwitchChatWorkflow, input).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	err = workflow.ExecuteChildWorkflow(ctx, MoveTwitchChatWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func ConvertTwitchLiveChatWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
			MaximumInterval:    15 * time.Minute,
		},
	})

	err := workflow.ExecuteActivity(ctx, activities.ConvertTwitchLiveChat, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "convert-chat")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil

}

// *Mid Level Workflow*
func ArchiveTwitchChatWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {

	err := workflow.ExecuteChildWorkflow(ctx, DownloadTwitchChatWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	if input.Queue.RenderChat {
		err = workflow.ExecuteChildWorkflow(ctx, RenderTwitchChatWorkflow, input).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	err = workflow.ExecuteChildWorkflow(ctx, MoveTwitchChatWorkflow, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func DownloadTwitchVideoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	cctx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "video-download",
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
			MaximumInterval:    15 * time.Minute,
		},
	})

	err := workflow.ExecuteActivity(cctx, activities.DownloadTwitchVideo, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "download-video")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func DownloadTwitchLiveVideoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    1,
			MaximumInterval:    15 * time.Minute,
		},
	})

	err := workflow.ExecuteActivity(ctx, activities.DownloadTwitchLiveVideo, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "download-video")
	}

	future := workflow.RequestCancelExternalWorkflow(ctx, input.LiveChatWorkflowId, "")
	if err := future.Get(ctx, nil); err != nil {
		return err
	}

	// mark live channel as not live
	live, err := database.DB().Client.Live.Query().Where(live.ID(input.LiveWatchChannel.ID)).Only(context.Background())
	if err != nil {
		// allow not found error to pass
		if _, ok := err.(*ent.NotFoundError); !ok {
			log.Error().Err(err).Msg("error getting live channel")
			return err
		}
	}
	if live != nil {
		_, err = live.Update().SetIsLive(false).Save(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("error updating live channel")
			return err
		}
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func PostprocessVideoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "video-convert",
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
			MaximumInterval:    15 * time.Minute,
		},
	})

	err := workflow.ExecuteActivity(ctx, activities.PostprocessVideo, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "postprocess-video")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func MoveVideoWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
			MaximumInterval:    15 * time.Minute,
		},
	})

	err := workflow.ExecuteActivity(ctx, activities.MoveVideo, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "move-video")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func DownloadTwitchChatWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "chat-download",
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
			MaximumInterval:    15 * time.Minute,
		},
	})

	err := workflow.ExecuteActivity(ctx, activities.DownloadTwitchChat, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "download-chat")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func DownloadTwitchLiveChatWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    1,
			MaximumInterval:    15 * time.Minute,
		},
		WaitForCancellation: false,
	})

	defer func() {

		if !errors.Is(ctx.Err(), workflow.ErrCanceled) {
			return
		}

		// When the Workflow is canceled, it has to get a new disconnected context to execute any Activities
		log.Debug().Msgf("Killing chat download: %s", input.LiveChatWorkflowId)
		newCtx, _ := workflow.NewDisconnectedContext(ctx)
		err := workflow.ExecuteActivity(newCtx, activities.KillTwitchLiveChatDownload, input).Get(ctx, nil)
		if err != nil {
			log.Error().Err(err).Msgf("error killing chat download: %v", err)
		}

		log.Debug().Msgf("Sending signal to continue chat archive: %s", input.LiveChatArchiveWorkflowId)
		signal := utils.ArchiveTwitchLiveChatStartSignal{
			Start: true,
		}
		err = workflow.SignalExternalWorkflow(ctx, input.LiveChatArchiveWorkflowId, "", "continue-chat-arhive", signal).Get(ctx, nil)
		if err != nil {
			log.Error().Err(err).Msgf("error sending signal to continue chat archive: %v", err)
		}
	}()

	var signal utils.ArchiveTwitchLiveChatStartSignal
	signalChan := workflow.GetSignalChannel(ctx, "start-chat-download")
	signalChan.Receive(ctx, &signal)

	log.Info().Msgf("Received signal: %v", signal)

	err := workflow.ExecuteActivity(ctx, activities.DownloadTwitchLiveChat, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func RenderTwitchChatWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "chat-render",
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
			MaximumInterval:    15 * time.Minute,
		},
	})

	err := workflow.ExecuteActivity(ctx, activities.RenderTwitchChat, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "render-chat")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func MoveTwitchChatWorkflow(ctx workflow.Context, input dto.ArchiveVideoInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
			MaximumInterval:    15 * time.Minute,
		},
	})

	err := workflow.ExecuteActivity(ctx, activities.MoveChat, input).Get(ctx, nil)
	if err != nil {
		return workflowErrorHandler(err, input, "move-chat")
	}

	err = checkIfTasksAreDone(input)
	if err != nil {
		return err
	}

	return nil
}

// *Low Level Workflow*
func SaveTwitchVideoChapters(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		HeartbeatTimeout:    90 * time.Second,
		StartToCloseTimeout: 168 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    3,
			MaximumInterval:    15 * time.Minute,
		},
	})

	err := workflow.ExecuteActivity(ctx, activities.TwitchSaveVideoChapters).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
