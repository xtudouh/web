CREATE TABLE if not EXISTS p_activity (
  id            BIGINT PRIMARY KEY,
  longitude     double precision,
  latitude      double precision,
  address       varchar(100),
  activityname  varchar(200),
  brief         varchar(1000),
  date          TIMESTAMP,
  crt           TIMESTAMP,
  lut           TIMESTAMP,
  del           boolean
);

COMMENT ON TABLE p_activity IS '活动表';
COMMENT ON COLUMN p_activity.id IS '主键';
COMMENT ON COLUMN p_activity.address IS '地址';
COMMENT ON COLUMN p_activity.activityname IS '活动名称';
COMMENT ON COLUMN p_activity.longitude IS '经度';
COMMENT ON COLUMN p_activity.latitude IS '纬度';
COMMENT ON COLUMN p_activity.date IS '举办时间';
COMMENT ON COLUMN p_activity.crt IS '创建时间';
COMMENT ON COLUMN p_activity.lut IS '最后更新时间';
COMMENT ON COLUMN p_activity.del IS '是否删除';

CREATE TABLE if not EXISTS p_signin (
  id            BIGINT PRIMARY KEY,
  activityid    BIGINT,
  custname      varchar(100),
  mobile        varchar(20),
  crt           TIMESTAMP,
  lut           TIMESTAMP,
  UNIQUE (activityid, mobile)
);

COMMENT ON TABLE p_signin IS '签到表';
COMMENT ON COLUMN p_signin.id IS '主键';
COMMENT ON COLUMN p_signin.activityid IS '活动ID';
COMMENT ON COLUMN p_signin.custname IS '客户名称';
COMMENT ON COLUMN p_signin.mobile IS '手机';
COMMENT ON COLUMN p_signin.crt IS '创建时间';
COMMENT ON COLUMN p_signin.lut IS '最后更新时间';