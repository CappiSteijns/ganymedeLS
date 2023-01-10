// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/queue"
	"github.com/zibbp/ganymede/ent/vod"
	"github.com/zibbp/ganymede/internal/utils"
)

// QueueCreate is the builder for creating a Queue entity.
type QueueCreate struct {
	config
	mutation *QueueMutation
	hooks    []Hook
}

// SetLiveArchive sets the "live_archive" field.
func (qc *QueueCreate) SetLiveArchive(b bool) *QueueCreate {
	qc.mutation.SetLiveArchive(b)
	return qc
}

// SetNillableLiveArchive sets the "live_archive" field if the given value is not nil.
func (qc *QueueCreate) SetNillableLiveArchive(b *bool) *QueueCreate {
	if b != nil {
		qc.SetLiveArchive(*b)
	}
	return qc
}

// SetOnHold sets the "on_hold" field.
func (qc *QueueCreate) SetOnHold(b bool) *QueueCreate {
	qc.mutation.SetOnHold(b)
	return qc
}

// SetNillableOnHold sets the "on_hold" field if the given value is not nil.
func (qc *QueueCreate) SetNillableOnHold(b *bool) *QueueCreate {
	if b != nil {
		qc.SetOnHold(*b)
	}
	return qc
}

// SetVideoProcessing sets the "video_processing" field.
func (qc *QueueCreate) SetVideoProcessing(b bool) *QueueCreate {
	qc.mutation.SetVideoProcessing(b)
	return qc
}

// SetNillableVideoProcessing sets the "video_processing" field if the given value is not nil.
func (qc *QueueCreate) SetNillableVideoProcessing(b *bool) *QueueCreate {
	if b != nil {
		qc.SetVideoProcessing(*b)
	}
	return qc
}

// SetChatProcessing sets the "chat_processing" field.
func (qc *QueueCreate) SetChatProcessing(b bool) *QueueCreate {
	qc.mutation.SetChatProcessing(b)
	return qc
}

// SetNillableChatProcessing sets the "chat_processing" field if the given value is not nil.
func (qc *QueueCreate) SetNillableChatProcessing(b *bool) *QueueCreate {
	if b != nil {
		qc.SetChatProcessing(*b)
	}
	return qc
}

// SetProcessing sets the "processing" field.
func (qc *QueueCreate) SetProcessing(b bool) *QueueCreate {
	qc.mutation.SetProcessing(b)
	return qc
}

// SetNillableProcessing sets the "processing" field if the given value is not nil.
func (qc *QueueCreate) SetNillableProcessing(b *bool) *QueueCreate {
	if b != nil {
		qc.SetProcessing(*b)
	}
	return qc
}

// SetTaskVodCreateFolder sets the "task_vod_create_folder" field.
func (qc *QueueCreate) SetTaskVodCreateFolder(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskVodCreateFolder(us)
	return qc
}

// SetNillableTaskVodCreateFolder sets the "task_vod_create_folder" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskVodCreateFolder(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskVodCreateFolder(*us)
	}
	return qc
}

// SetTaskVodDownloadThumbnail sets the "task_vod_download_thumbnail" field.
func (qc *QueueCreate) SetTaskVodDownloadThumbnail(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskVodDownloadThumbnail(us)
	return qc
}

// SetNillableTaskVodDownloadThumbnail sets the "task_vod_download_thumbnail" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskVodDownloadThumbnail(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskVodDownloadThumbnail(*us)
	}
	return qc
}

// SetTaskVodSaveInfo sets the "task_vod_save_info" field.
func (qc *QueueCreate) SetTaskVodSaveInfo(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskVodSaveInfo(us)
	return qc
}

// SetNillableTaskVodSaveInfo sets the "task_vod_save_info" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskVodSaveInfo(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskVodSaveInfo(*us)
	}
	return qc
}

// SetTaskVideoDownload sets the "task_video_download" field.
func (qc *QueueCreate) SetTaskVideoDownload(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskVideoDownload(us)
	return qc
}

// SetNillableTaskVideoDownload sets the "task_video_download" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskVideoDownload(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskVideoDownload(*us)
	}
	return qc
}

// SetTaskVideoConvert sets the "task_video_convert" field.
func (qc *QueueCreate) SetTaskVideoConvert(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskVideoConvert(us)
	return qc
}

// SetNillableTaskVideoConvert sets the "task_video_convert" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskVideoConvert(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskVideoConvert(*us)
	}
	return qc
}

// SetTaskVideoMove sets the "task_video_move" field.
func (qc *QueueCreate) SetTaskVideoMove(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskVideoMove(us)
	return qc
}

// SetNillableTaskVideoMove sets the "task_video_move" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskVideoMove(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskVideoMove(*us)
	}
	return qc
}

// SetTaskChatDownload sets the "task_chat_download" field.
func (qc *QueueCreate) SetTaskChatDownload(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskChatDownload(us)
	return qc
}

// SetNillableTaskChatDownload sets the "task_chat_download" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskChatDownload(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskChatDownload(*us)
	}
	return qc
}

// SetTaskChatConvert sets the "task_chat_convert" field.
func (qc *QueueCreate) SetTaskChatConvert(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskChatConvert(us)
	return qc
}

// SetNillableTaskChatConvert sets the "task_chat_convert" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskChatConvert(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskChatConvert(*us)
	}
	return qc
}

// SetTaskChatRender sets the "task_chat_render" field.
func (qc *QueueCreate) SetTaskChatRender(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskChatRender(us)
	return qc
}

// SetNillableTaskChatRender sets the "task_chat_render" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskChatRender(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskChatRender(*us)
	}
	return qc
}

// SetTaskChatMove sets the "task_chat_move" field.
func (qc *QueueCreate) SetTaskChatMove(us utils.TaskStatus) *QueueCreate {
	qc.mutation.SetTaskChatMove(us)
	return qc
}

// SetNillableTaskChatMove sets the "task_chat_move" field if the given value is not nil.
func (qc *QueueCreate) SetNillableTaskChatMove(us *utils.TaskStatus) *QueueCreate {
	if us != nil {
		qc.SetTaskChatMove(*us)
	}
	return qc
}

// SetChatStart sets the "chat_start" field.
func (qc *QueueCreate) SetChatStart(t time.Time) *QueueCreate {
	qc.mutation.SetChatStart(t)
	return qc
}

// SetNillableChatStart sets the "chat_start" field if the given value is not nil.
func (qc *QueueCreate) SetNillableChatStart(t *time.Time) *QueueCreate {
	if t != nil {
		qc.SetChatStart(*t)
	}
	return qc
}

// SetUpdatedAt sets the "updated_at" field.
func (qc *QueueCreate) SetUpdatedAt(t time.Time) *QueueCreate {
	qc.mutation.SetUpdatedAt(t)
	return qc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (qc *QueueCreate) SetNillableUpdatedAt(t *time.Time) *QueueCreate {
	if t != nil {
		qc.SetUpdatedAt(*t)
	}
	return qc
}

// SetCreatedAt sets the "created_at" field.
func (qc *QueueCreate) SetCreatedAt(t time.Time) *QueueCreate {
	qc.mutation.SetCreatedAt(t)
	return qc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (qc *QueueCreate) SetNillableCreatedAt(t *time.Time) *QueueCreate {
	if t != nil {
		qc.SetCreatedAt(*t)
	}
	return qc
}

// SetID sets the "id" field.
func (qc *QueueCreate) SetID(u uuid.UUID) *QueueCreate {
	qc.mutation.SetID(u)
	return qc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (qc *QueueCreate) SetNillableID(u *uuid.UUID) *QueueCreate {
	if u != nil {
		qc.SetID(*u)
	}
	return qc
}

// SetVodID sets the "vod" edge to the Vod entity by ID.
func (qc *QueueCreate) SetVodID(id uuid.UUID) *QueueCreate {
	qc.mutation.SetVodID(id)
	return qc
}

// SetVod sets the "vod" edge to the Vod entity.
func (qc *QueueCreate) SetVod(v *Vod) *QueueCreate {
	return qc.SetVodID(v.ID)
}

// Mutation returns the QueueMutation object of the builder.
func (qc *QueueCreate) Mutation() *QueueMutation {
	return qc.mutation
}

// Save creates the Queue in the database.
func (qc *QueueCreate) Save(ctx context.Context) (*Queue, error) {
	qc.defaults()
	return withHooks[*Queue, QueueMutation](ctx, qc.sqlSave, qc.mutation, qc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QueueCreate) SaveX(ctx context.Context) *Queue {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qc *QueueCreate) Exec(ctx context.Context) error {
	_, err := qc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qc *QueueCreate) ExecX(ctx context.Context) {
	if err := qc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qc *QueueCreate) defaults() {
	if _, ok := qc.mutation.LiveArchive(); !ok {
		v := queue.DefaultLiveArchive
		qc.mutation.SetLiveArchive(v)
	}
	if _, ok := qc.mutation.OnHold(); !ok {
		v := queue.DefaultOnHold
		qc.mutation.SetOnHold(v)
	}
	if _, ok := qc.mutation.VideoProcessing(); !ok {
		v := queue.DefaultVideoProcessing
		qc.mutation.SetVideoProcessing(v)
	}
	if _, ok := qc.mutation.ChatProcessing(); !ok {
		v := queue.DefaultChatProcessing
		qc.mutation.SetChatProcessing(v)
	}
	if _, ok := qc.mutation.Processing(); !ok {
		v := queue.DefaultProcessing
		qc.mutation.SetProcessing(v)
	}
	if _, ok := qc.mutation.TaskVodCreateFolder(); !ok {
		v := queue.DefaultTaskVodCreateFolder
		qc.mutation.SetTaskVodCreateFolder(v)
	}
	if _, ok := qc.mutation.TaskVodDownloadThumbnail(); !ok {
		v := queue.DefaultTaskVodDownloadThumbnail
		qc.mutation.SetTaskVodDownloadThumbnail(v)
	}
	if _, ok := qc.mutation.TaskVodSaveInfo(); !ok {
		v := queue.DefaultTaskVodSaveInfo
		qc.mutation.SetTaskVodSaveInfo(v)
	}
	if _, ok := qc.mutation.TaskVideoDownload(); !ok {
		v := queue.DefaultTaskVideoDownload
		qc.mutation.SetTaskVideoDownload(v)
	}
	if _, ok := qc.mutation.TaskVideoConvert(); !ok {
		v := queue.DefaultTaskVideoConvert
		qc.mutation.SetTaskVideoConvert(v)
	}
	if _, ok := qc.mutation.TaskVideoMove(); !ok {
		v := queue.DefaultTaskVideoMove
		qc.mutation.SetTaskVideoMove(v)
	}
	if _, ok := qc.mutation.TaskChatDownload(); !ok {
		v := queue.DefaultTaskChatDownload
		qc.mutation.SetTaskChatDownload(v)
	}
	if _, ok := qc.mutation.TaskChatConvert(); !ok {
		v := queue.DefaultTaskChatConvert
		qc.mutation.SetTaskChatConvert(v)
	}
	if _, ok := qc.mutation.TaskChatRender(); !ok {
		v := queue.DefaultTaskChatRender
		qc.mutation.SetTaskChatRender(v)
	}
	if _, ok := qc.mutation.TaskChatMove(); !ok {
		v := queue.DefaultTaskChatMove
		qc.mutation.SetTaskChatMove(v)
	}
	if _, ok := qc.mutation.UpdatedAt(); !ok {
		v := queue.DefaultUpdatedAt()
		qc.mutation.SetUpdatedAt(v)
	}
	if _, ok := qc.mutation.CreatedAt(); !ok {
		v := queue.DefaultCreatedAt()
		qc.mutation.SetCreatedAt(v)
	}
	if _, ok := qc.mutation.ID(); !ok {
		v := queue.DefaultID()
		qc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qc *QueueCreate) check() error {
	if _, ok := qc.mutation.LiveArchive(); !ok {
		return &ValidationError{Name: "live_archive", err: errors.New(`ent: missing required field "Queue.live_archive"`)}
	}
	if _, ok := qc.mutation.OnHold(); !ok {
		return &ValidationError{Name: "on_hold", err: errors.New(`ent: missing required field "Queue.on_hold"`)}
	}
	if _, ok := qc.mutation.VideoProcessing(); !ok {
		return &ValidationError{Name: "video_processing", err: errors.New(`ent: missing required field "Queue.video_processing"`)}
	}
	if _, ok := qc.mutation.ChatProcessing(); !ok {
		return &ValidationError{Name: "chat_processing", err: errors.New(`ent: missing required field "Queue.chat_processing"`)}
	}
	if _, ok := qc.mutation.Processing(); !ok {
		return &ValidationError{Name: "processing", err: errors.New(`ent: missing required field "Queue.processing"`)}
	}
	if v, ok := qc.mutation.TaskVodCreateFolder(); ok {
		if err := queue.TaskVodCreateFolderValidator(v); err != nil {
			return &ValidationError{Name: "task_vod_create_folder", err: fmt.Errorf(`ent: validator failed for field "Queue.task_vod_create_folder": %w`, err)}
		}
	}
	if v, ok := qc.mutation.TaskVodDownloadThumbnail(); ok {
		if err := queue.TaskVodDownloadThumbnailValidator(v); err != nil {
			return &ValidationError{Name: "task_vod_download_thumbnail", err: fmt.Errorf(`ent: validator failed for field "Queue.task_vod_download_thumbnail": %w`, err)}
		}
	}
	if v, ok := qc.mutation.TaskVodSaveInfo(); ok {
		if err := queue.TaskVodSaveInfoValidator(v); err != nil {
			return &ValidationError{Name: "task_vod_save_info", err: fmt.Errorf(`ent: validator failed for field "Queue.task_vod_save_info": %w`, err)}
		}
	}
	if v, ok := qc.mutation.TaskVideoDownload(); ok {
		if err := queue.TaskVideoDownloadValidator(v); err != nil {
			return &ValidationError{Name: "task_video_download", err: fmt.Errorf(`ent: validator failed for field "Queue.task_video_download": %w`, err)}
		}
	}
	if v, ok := qc.mutation.TaskVideoConvert(); ok {
		if err := queue.TaskVideoConvertValidator(v); err != nil {
			return &ValidationError{Name: "task_video_convert", err: fmt.Errorf(`ent: validator failed for field "Queue.task_video_convert": %w`, err)}
		}
	}
	if v, ok := qc.mutation.TaskVideoMove(); ok {
		if err := queue.TaskVideoMoveValidator(v); err != nil {
			return &ValidationError{Name: "task_video_move", err: fmt.Errorf(`ent: validator failed for field "Queue.task_video_move": %w`, err)}
		}
	}
	if v, ok := qc.mutation.TaskChatDownload(); ok {
		if err := queue.TaskChatDownloadValidator(v); err != nil {
			return &ValidationError{Name: "task_chat_download", err: fmt.Errorf(`ent: validator failed for field "Queue.task_chat_download": %w`, err)}
		}
	}
	if v, ok := qc.mutation.TaskChatConvert(); ok {
		if err := queue.TaskChatConvertValidator(v); err != nil {
			return &ValidationError{Name: "task_chat_convert", err: fmt.Errorf(`ent: validator failed for field "Queue.task_chat_convert": %w`, err)}
		}
	}
	if v, ok := qc.mutation.TaskChatRender(); ok {
		if err := queue.TaskChatRenderValidator(v); err != nil {
			return &ValidationError{Name: "task_chat_render", err: fmt.Errorf(`ent: validator failed for field "Queue.task_chat_render": %w`, err)}
		}
	}
	if v, ok := qc.mutation.TaskChatMove(); ok {
		if err := queue.TaskChatMoveValidator(v); err != nil {
			return &ValidationError{Name: "task_chat_move", err: fmt.Errorf(`ent: validator failed for field "Queue.task_chat_move": %w`, err)}
		}
	}
	if _, ok := qc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Queue.updated_at"`)}
	}
	if _, ok := qc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Queue.created_at"`)}
	}
	if _, ok := qc.mutation.VodID(); !ok {
		return &ValidationError{Name: "vod", err: errors.New(`ent: missing required edge "Queue.vod"`)}
	}
	return nil
}

func (qc *QueueCreate) sqlSave(ctx context.Context) (*Queue, error) {
	if err := qc.check(); err != nil {
		return nil, err
	}
	_node, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	qc.mutation.id = &_node.ID
	qc.mutation.done = true
	return _node, nil
}

func (qc *QueueCreate) createSpec() (*Queue, *sqlgraph.CreateSpec) {
	var (
		_node = &Queue{config: qc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: queue.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: queue.FieldID,
			},
		}
	)
	if id, ok := qc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := qc.mutation.LiveArchive(); ok {
		_spec.SetField(queue.FieldLiveArchive, field.TypeBool, value)
		_node.LiveArchive = value
	}
	if value, ok := qc.mutation.OnHold(); ok {
		_spec.SetField(queue.FieldOnHold, field.TypeBool, value)
		_node.OnHold = value
	}
	if value, ok := qc.mutation.VideoProcessing(); ok {
		_spec.SetField(queue.FieldVideoProcessing, field.TypeBool, value)
		_node.VideoProcessing = value
	}
	if value, ok := qc.mutation.ChatProcessing(); ok {
		_spec.SetField(queue.FieldChatProcessing, field.TypeBool, value)
		_node.ChatProcessing = value
	}
	if value, ok := qc.mutation.Processing(); ok {
		_spec.SetField(queue.FieldProcessing, field.TypeBool, value)
		_node.Processing = value
	}
	if value, ok := qc.mutation.TaskVodCreateFolder(); ok {
		_spec.SetField(queue.FieldTaskVodCreateFolder, field.TypeEnum, value)
		_node.TaskVodCreateFolder = value
	}
	if value, ok := qc.mutation.TaskVodDownloadThumbnail(); ok {
		_spec.SetField(queue.FieldTaskVodDownloadThumbnail, field.TypeEnum, value)
		_node.TaskVodDownloadThumbnail = value
	}
	if value, ok := qc.mutation.TaskVodSaveInfo(); ok {
		_spec.SetField(queue.FieldTaskVodSaveInfo, field.TypeEnum, value)
		_node.TaskVodSaveInfo = value
	}
	if value, ok := qc.mutation.TaskVideoDownload(); ok {
		_spec.SetField(queue.FieldTaskVideoDownload, field.TypeEnum, value)
		_node.TaskVideoDownload = value
	}
	if value, ok := qc.mutation.TaskVideoConvert(); ok {
		_spec.SetField(queue.FieldTaskVideoConvert, field.TypeEnum, value)
		_node.TaskVideoConvert = value
	}
	if value, ok := qc.mutation.TaskVideoMove(); ok {
		_spec.SetField(queue.FieldTaskVideoMove, field.TypeEnum, value)
		_node.TaskVideoMove = value
	}
	if value, ok := qc.mutation.TaskChatDownload(); ok {
		_spec.SetField(queue.FieldTaskChatDownload, field.TypeEnum, value)
		_node.TaskChatDownload = value
	}
	if value, ok := qc.mutation.TaskChatConvert(); ok {
		_spec.SetField(queue.FieldTaskChatConvert, field.TypeEnum, value)
		_node.TaskChatConvert = value
	}
	if value, ok := qc.mutation.TaskChatRender(); ok {
		_spec.SetField(queue.FieldTaskChatRender, field.TypeEnum, value)
		_node.TaskChatRender = value
	}
	if value, ok := qc.mutation.TaskChatMove(); ok {
		_spec.SetField(queue.FieldTaskChatMove, field.TypeEnum, value)
		_node.TaskChatMove = value
	}
	if value, ok := qc.mutation.ChatStart(); ok {
		_spec.SetField(queue.FieldChatStart, field.TypeTime, value)
		_node.ChatStart = value
	}
	if value, ok := qc.mutation.UpdatedAt(); ok {
		_spec.SetField(queue.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := qc.mutation.CreatedAt(); ok {
		_spec.SetField(queue.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := qc.mutation.VodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   queue.VodTable,
			Columns: []string{queue.VodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.vod_queue = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// QueueCreateBulk is the builder for creating many Queue entities in bulk.
type QueueCreateBulk struct {
	config
	builders []*QueueCreate
}

// Save creates the Queue entities in the database.
func (qcb *QueueCreateBulk) Save(ctx context.Context) ([]*Queue, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qcb.builders))
	nodes := make([]*Queue, len(qcb.builders))
	mutators := make([]Mutator, len(qcb.builders))
	for i := range qcb.builders {
		func(i int, root context.Context) {
			builder := qcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QueueMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, qcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, qcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qcb *QueueCreateBulk) SaveX(ctx context.Context) []*Queue {
	v, err := qcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qcb *QueueCreateBulk) Exec(ctx context.Context) error {
	_, err := qcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcb *QueueCreateBulk) ExecX(ctx context.Context) {
	if err := qcb.Exec(ctx); err != nil {
		panic(err)
	}
}
