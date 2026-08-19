# My Coffee Log Metrics Dictionary

> 文档版本：V1.0  
> 文档日期：2026-08-18  
> 目标：统一产品、前端、后端和数据团队对核心指标的口径定义。

---

## 1. 有效记录（Valid Record）

一条记录被认定为"有效记录"，必须同时满足以下条件：

- 由真实用户创建（非系统或测试账号）。
- `is_test_data = false`。
- 具备有效日期（`drink_date` 不为空）。
- 咖啡类型经过用户确认（`coffee_type` 非空）。

### 排除条件

- `is_test_data = true` 的记录。
- `drink_date` 为空的记录。
- `coffee_type` 为空的记录。

---

## 2. 有效感官记录（Valid Sensory Record）

在有效记录基础上，六维感官评分由用户主动提交：

- `sensory_recorded = true`。
- 六维值（acidity, bitterness, sweetness, body, aroma, aftertaste）均在 0–5 范围内。
- 不包含系统自动填充的默认 0 分。

### 用途

- 风味雷达平均值计算。
- 感官偏好统计。
- 咖啡人格规则匹配（需要感官数据的规则）。

---

## 3. 指标定义

### 3.1 获客与激活

| 指标 | 定义 | 计算公式 | 时间窗口 |
|---|---|---|---|
| `register_completed` | 用户完成注册 | 事件计数 | 按日/周/月 |
| `first_record_created` | 用户创建第一条有效记录 | COUNT(valid_records) WHERE record_sequence = 1 | 按用户 |
| `register_to_first_record_rate` | 注册后完成首条记录的比例 | first_record_created / register_completed | 按周/月 |
| `first_record_duration` | 首条记录从进入到保存的耗时 | record_created.occurred_at - record_page_viewed.occurred_at | 按用户 |

### 3.2 使用与留存

| 指标 | 定义 | 计算公式 | 时间窗口 |
|---|---|---|---|
| `DAU` | 日活跃用户 | COUNT(DISTINCT user_id) WHERE occurred_at >= today | 日 |
| `WAU` | 周活跃用户 | COUNT(DISTINCT user_id) WHERE occurred_at >= 7 days ago | 周 |
| `MAU` | 月活跃用户 | COUNT(DISTINCT user_id) WHERE occurred_at >= 30 days ago | 月 |
| `D1_retention` | 注册次日回访 | 用户在注册后第 1 天有至少一次页面访问 | 按日 |
| `D7_retention` | 注册 7 日回访 | 用户在注册后第 7 天有至少一次页面访问 | 按周 |
| `weekly_active_record_rate` | 周活跃中有记录的用户比例 | users_with_valid_record_this_wau / WAU | 周 |
| `avg_valid_records_per_week` | 每周人均有效记录数 | SUM(valid_records_this_week) / WAU | 周 |

### 3.3 功能价值

| 指标 | 定义 | 触发条件 |
|---|---|---|
| `quick_record_usage` | 快捷记录使用率 | record_mode = 'quick' 的创建次数 / 总创建次数 |
| `detailed_record_usage` | 精细记录使用率 | record_mode = 'detailed' 的创建次数 / 总创建次数 |
| `rebrew_usage` | 再次冲煮使用率 | source_log_id IS NOT NULL 的创建次数 / 总创建次数 |
| `weekly_review_view_rate` | 周度回顾查看率 | review_viewed(period=weekly) / WAU |
| `monthly_review_view_rate` | 月度回顾查看率 | review_viewed(period=monthly) / MAU |
| `share_export_rate` | 分享卡导出率 | share_exported 次数 / valid_records |
| `ai_content_keep_rate` | AI 内容保留率 | records where AI summary is present AND not edited / records with AI summary generated |

### 3.4 数据质量

| 指标 | 定义 | 计算公式 |
|---|---|---|
| `valid_sensory_ratio` | 有效感官评分占比 | COUNT(sensory_recorded=true) / COUNT(valid_records) |
| `valid_shop_ratio` | 有效门店数据占比 | COUNT(shop_source='user_input') / COUNT(valid_records) |
| `user_notes_ratio` | 用户原始笔记占比 | COUNT(notes_source='user_input') / COUNT(valid_records) |
| `system_suggestion_confirm_rate` | 系统建议确认率 | system_suggested accepted / system_suggested shown |
| `default_pollution_ratio` | 已识别默认污染数据比例 | COUNT(known_defaults) / total_records |

---

## 4. 阈值定义

| 能力 | 最低有效样本数 | 说明 |
|---|---:|---|
| 基础类型偏好 | 3 | 至少 3 条有效记录 |
| 初步风味雷达 | 3 | 至少 3 条有效感官记录 |
| 周度趋势 | 5 | 至少 5 条有效记录或覆盖 2 周 |
| 咖啡人格 | 8 | 至少 8 条有效记录 |
| 月度变化对比 | 6 | 连续 2 个月且每月不少于 3 条 |
| 历史记忆 | 1 | 至少 1 条符合条件的历史记录 |

---

## 5. 数据来源标记

| 标记 | 含义 | 是否参与统计 |
|---|---|---|
| `user_input` | 用户主动填写或确认 | 是 |
| `system_suggested` | 系统建议但尚未确认 | 否（除非用户确认） |
| `ai_generated` | AI 生成 | 否（不作为用户原始数据） |
| `recent_reuse` | 从历史记录复用 | 是 |
| `empty` | 用户未提供 | 否 |
