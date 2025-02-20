// Code generated by goctl. DO NOT EDIT.

package coupon

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	couponsFieldNames          = builder.RawFieldNames(&Coupons{})
	couponsRows                = strings.Join(couponsFieldNames, ",")
	couponsRowsExpectAutoSet   = strings.Join(stringx.Remove(couponsFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	couponsRowsWithPlaceHolder = strings.Join(stringx.Remove(couponsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	couponsModel interface {
		Insert(ctx context.Context, data *Coupons) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Coupons, error)
		Update(ctx context.Context, data *Coupons) error
		Delete(ctx context.Context, id string) error
	}

	defaultCouponsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Coupons struct {
		Id             string    `db:"id"`              // 优惠券ID
		Name           string    `db:"name"`            // 券名称
		Type           int64     `db:"type"`            // 类型：1-满减 2-折扣 3-立减
		Value          int64     `db:"value"`           // 优惠值（根据类型：分/百分比）
		MinAmount      int64     `db:"min_amount"`      // 最低消费金额（分）
		StartTime      time.Time `db:"start_time"`      // 有效期开始
		EndTime        time.Time `db:"end_time"`        // 有效期结束
		Status         int64     `db:"status"`          // 状态：0-禁用 1-启用
		TotalCount     uint64    `db:"total_count"`     // 发行总量
		RemainingCount uint64    `db:"remaining_count"` // 剩余数量
		CreatedAt      time.Time `db:"created_at"`      // 创建时间
		UpdatedAt      time.Time `db:"updated_at"`      // 更新时间
	}
)

func newCouponsModel(conn sqlx.SqlConn) *defaultCouponsModel {
	return &defaultCouponsModel{
		conn:  conn,
		table: "`coupons`",
	}
}

func (m *defaultCouponsModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCouponsModel) FindOne(ctx context.Context, id string) (*Coupons, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", couponsRows, m.table)
	var resp Coupons
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCouponsModel) Insert(ctx context.Context, data *Coupons) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, couponsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.Type, data.Value, data.MinAmount, data.StartTime, data.EndTime, data.Status, data.TotalCount, data.RemainingCount)
	return ret, err
}

func (m *defaultCouponsModel) Update(ctx context.Context, data *Coupons) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, couponsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Type, data.Value, data.MinAmount, data.StartTime, data.EndTime, data.Status, data.TotalCount, data.RemainingCount, data.Id)
	return err
}

func (m *defaultCouponsModel) tableName() string {
	return m.table
}
