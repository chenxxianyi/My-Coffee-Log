# My Coffee Log 项目开发方案

## 1. 项目概述

**My Coffee Log** 是一款面向日常咖啡爱好者的极简咖啡生活记录 App。

它不是专业咖啡参数工具，而是一款偏生活方式、审美记录、社交分享的咖啡手账产品。

核心体验是：

> 用户用 3 秒记录一杯咖啡，系统自动生成高级风味卡片，并沉淀成自己的咖啡生活时间线。

产品最终应让用户感觉：

> 我不是在记录咖啡，而是在记录我的生活方式。

---

## 2. 产品定位

### 2.1 目标用户

- 日常咖啡爱好者
- 喜欢记录生活方式的人
- 喜欢拍照、分享、收藏咖啡体验的人
- 不需要复杂专业参数，但希望记录味道、情绪和场景的人

### 2.2 产品关键词

- Minimal
- Editorial
- Premium
- Warm
- Calm
- Coffee Lifestyle
- Soft
- Breathable
- Typography Driven
- Magazine Layout

### 2.3 核心价值

- 快速记录每日咖啡
- 自动生成风味总结
- 沉淀咖啡生活时间线
- 生成高审美分享卡片
- 用轻量统计理解自己的咖啡偏好

---

## 3. 技术栈

### 3.1 前端技术栈

- Vue 3
- Vite
- TypeScript
- Pinia
- Vue Router
- Tailwind CSS
- ECharts 或 SVG 自定义雷达图
- html2canvas
- Capacitor

### 3.2 后端技术栈

- Go 1.22+
- Gin
- GORM
- MySQL 8
- Redis
- JWT
- Zap 日志
- Viper 配置
- Validator 参数校验
- Swaggo Swagger 文档
- Docker Compose

### 3.3 AI 能力

MVP 阶段先做本地 Mock AI 风味总结。

后续预留 OpenAI API 接入能力，支持以下环境变量：

- `OPENAI_API_KEY`
- `OPENAI_BASE_URL`
- `OPENAI_MODEL`

---

## 4. 整体架构设计

项目采用前后端分离架构。

```text
My Coffee Log
├── frontend/                 # Vue 3 前端应用
├── backend/                  # Go Gin 后端服务
├── docker-compose.yml        # 本地一键启动 MySQL、Redis、后端、前端
├── README.md                 # 项目说明
└── PROJECT_DEVELOPMENT_PLAN.md
```

系统主要由以下模块组成：

- 用户认证模块
- 用户资料模块
- 咖啡记录模块
- 风味标签模块
- AI 风味总结模块
- 时间线模块
- 数据统计模块
- 分享卡片模块

---

## 5. 后端架构方案

后端采用清晰分层架构，避免 Handler 直接访问数据库。

```text
backend/
  cmd/
    server/
      main.go
  internal/
    config/
      config.go
    database/
      mysql.go
      redis.go
    middleware/
      auth.go
      cors.go
      logger.go
      recovery.go
    model/
      user.go
      coffee_log.go
      flavor_tag.go
      card_template.go
    repository/
      user_repository.go
      coffee_log_repository.go
      flavor_tag_repository.go
      stats_repository.go
    service/
      auth_service.go
      user_service.go
      coffee_log_service.go
      stats_service.go
      ai_service.go
    handler/
      auth_handler.go
      user_handler.go
      coffee_log_handler.go
      stats_handler.go
      ai_handler.go
    router/
      router.go
    response/
      response.go
    utils/
      jwt.go
      password.go
      pagination.go
      validator.go
  migrations/
  docs/
  go.mod
  go.sum
  Dockerfile
  .env.example
```

### 5.1 分层职责

- **Handler 层**：处理 HTTP 请求、参数绑定、调用 Service、返回统一响应。
- **Service 层**：处理业务逻辑，如注册、登录、创建咖啡记录、AI 总结生成。
- **Repository 层**：封装数据库访问逻辑。
- **Model 层**：定义 GORM 模型和关联关系。
- **Middleware 层**：处理 JWT、CORS、日志、异常恢复。
- **Response 层**：统一 API 返回格式。
- **Utils 层**：封装 JWT、密码加密、分页、参数校验等工具。

---

## 6. 前端架构方案

```text
frontend/
  src/
    api/
      auth.ts
      coffeeLog.ts
      stats.ts
      ai.ts
    assets/
    components/
      base/
      coffee/
      charts/
      share/
    layouts/
    pages/
      Home.vue
      Login.vue
      Register.vue
      CreateCoffeeLog.vue
      CoffeeDetail.vue
      Timeline.vue
      Stats.vue
      Profile.vue
    router/
    stores/
      auth.ts
      coffeeLog.ts
    styles/
    utils/
    App.vue
    main.ts
```

### 6.1 前端职责拆分

- **api/**：封装后端接口请求。
- **stores/**：管理登录状态、用户信息、咖啡记录状态。
- **pages/**：页面级组件。
- **components/base/**：按钮、输入框、标签、弹窗等基础组件。
- **components/coffee/**：咖啡记录卡片、创建步骤、标签选择等业务组件。
- **components/charts/**：风味雷达图。
- **components/share/**：分享卡片预览和导出。
- **layouts/**：登录页布局、主应用布局。
- **styles/**：全局样式、Tailwind 扩展、设计变量。

---

## 7. 核心功能模块

## 7.1 用户系统

### 功能范围

- 用户注册
- 用户登录
- JWT 鉴权
- 获取当前用户信息
- 修改昵称
- 修改头像

### 用户字段

- `id`
- `email`
- `nickname`
- `password_hash`
- `avatar_url`
- `created_at`
- `updated_at`

### 接口

- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `GET /api/v1/users/me`
- `PUT /api/v1/users/me`

### 实现要点

- 密码使用 bcrypt 加密。
- 登录成功后返回 JWT。
- 注册和登录以外的接口全部需要鉴权。
- JWT 中间件解析用户 ID 并写入请求上下文。

---

## 7.2 Coffee Log 咖啡记录

### 功能范围

用户可以创建、查看、编辑、删除自己的每日咖啡记录。

### 字段

- `coffee_name`
- `coffee_type`
- `shop_name`
- `location`
- `image_url`
- `drink_date`
- `mood`
- `notes`
- `acidity`
- `bitterness`
- `sweetness`
- `body`
- `aroma`
- `aftertaste`
- `ai_summary`

### 咖啡类型

- Americano
- Latte
- Pour Over
- Cold Brew
- Espresso
- Dirty
- Cappuccino
- Other

### 风味维度

- `acidity` 酸度 0-5
- `bitterness` 苦感 0-5
- `sweetness` 甜感 0-5
- `body` 醇厚度 0-5
- `aroma` 香气 0-5
- `aftertaste` 余韵 0-5

### 接口

- `POST /api/v1/coffee-logs`
- `GET /api/v1/coffee-logs`
- `GET /api/v1/coffee-logs/:id`
- `PUT /api/v1/coffee-logs/:id`
- `DELETE /api/v1/coffee-logs/:id`

### 筛选能力

列表接口需要支持：

- 分页
- 按月份筛选
- 按咖啡类型筛选
- 按风味标签筛选
- 默认按 `drink_date` 倒序排列

### 权限要求

用户只能查看、修改、删除自己的记录。

所有 Coffee Log 查询必须包含当前登录用户的 `user_id` 条件。

---

## 7.3 风味标签

### 默认标签

- floral 花香
- citrus 柑橘
- berry 莓果
- nutty 坚果
- chocolate 巧克力
- caramel 焦糖
- creamy 奶油
- winey 酒香
- smoky 烟熏
- herbal 草本

### 实现要点

- `CoffeeLog` 与 `FlavorTag` 是多对多关系。
- 系统初始化时写入默认风味标签。
- 创建和更新 Coffee Log 时支持绑定风味标签。
- Timeline、列表和统计需要支持按标签筛选。

---

## 7.4 Home 首页

首页不是普通 Dashboard，而应该像高级生活方式杂志封面。

### 展示内容

- 今日咖啡入口
- 最近一条咖啡记录
- 最近 7 天咖啡记录
- AI 今日风味摘要
- 快速创建按钮
- 本月咖啡次数

### 设计要求

- 大图优先
- 非对称排版
- 大量留白
- 杂志式标题层级
- 温暖、安静、高级

---

## 7.5 咖啡记录创建页

创建流程分三步。

### 第一步

- 选择咖啡类型
- 上传图片或填写图片地址
- 填写咖啡名称

### 第二步

- 使用滑动条选择六个风味强度

### 第三步

- 选择风味标签
- 选择心情
- 填写店铺
- 填写备注

### UX 原则

- 3 秒可完成最小记录。
- 非必填字段不阻塞用户。
- 操作轻量、有生活感。
- 提交后进入详情页。

---

## 7.6 咖啡详情页

### 展示内容

- 咖啡大图
- 咖啡名称
- 日期
- 店铺
- 风味雷达图
- 风味标签
- AI 风味描述
- 心情
- 备注
- 生成分享卡片按钮

### 设计要求

详情页应像一本生活方式杂志的单篇咖啡内页。

---

## 7.7 风味雷达图

### 数据来源

- `acidity`
- `bitterness`
- `sweetness`
- `body`
- `aroma`
- `aftertaste`

### 实现建议

优先使用 SVG 自定义雷达图。

原因：

- 风格更可控
- 更容易做成杂志感信息图
- 分享卡片导出更稳定
- 避免 ECharts 默认商业图表风格

### 视觉要求

- 简洁
- 柔和
- 低饱和
- 信息图感
- 不要电竞风
- 不要复杂商业图表风

---

## 7.8 AI 风味总结

### MVP 方案

后端通过本地 Mock 生成风味描述。

示例：

> 今天这杯咖啡带有柔和的柑橘酸质和轻微坚果尾韵，整体口感清爽，适合一个安静的下午。

### 接口

- `POST /api/v1/ai/flavor-summary`

### 实现要求

后端封装：

```text
internal/service/ai_service.go
```

AI Service 需要预留真实 AI 接口能力，后续可接入 OpenAI。

---

## 7.9 Timeline 时间线

### 功能范围

按日期展示用户所有咖啡记录。

### 筛选能力

- 按月份筛选
- 按咖啡类型筛选
- 按风味标签筛选

### 设计要求

- 生活方式时间流
- 图片优先
- 大量留白
- 像一本咖啡日记
- 不做普通后台列表

---

## 7.10 Coffee Share Card 分享卡片

分享卡片是项目核心功能之一。

### 卡片内容

- 咖啡图片
- 咖啡名称
- 日期
- 店铺
- 风味雷达图
- 风味标签
- AI 文案
- My Coffee Log 品牌标识

### 支持尺寸

- 1:1 小红书 / Instagram
- 3:4 小红书封面
- 9:16 Story

### 前端实现方案

- 使用 Vue 组件渲染卡片。
- 使用 html2canvas 导出图片。
- 支持保存到本地。

### 建议组件

- `ShareCardPreview.vue`
- `CoffeeShareCard.vue`
- `share.ts`

---

## 7.11 数据统计

### 功能范围

- 本月喝咖啡次数
- 总记录数
- 最常喝的咖啡类型
- 最常出现的风味标签
- 风味偏好平均值
- 按月份统计记录数量

### 接口

- `GET /api/v1/stats/overview`
- `GET /api/v1/stats/flavor-profile`
- `GET /api/v1/stats/monthly`

### 设计要求

统计页不应该像商业 BI 或后台报表，而应保持杂志感、柔和、轻量。

---

## 8. 数据库设计

### 8.1 users

```sql
CREATE TABLE users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  email VARCHAR(255) UNIQUE NOT NULL,
  nickname VARCHAR(100),
  password_hash VARCHAR(255) NOT NULL,
  avatar_url VARCHAR(500),
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME NULL
);
```

### 8.2 coffee_logs

```sql
CREATE TABLE coffee_logs (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  user_id BIGINT NOT NULL,
  coffee_name VARCHAR(255),
  coffee_type VARCHAR(50),
  shop_name VARCHAR(255),
  location VARCHAR(255),
  image_url VARCHAR(500),
  drink_date DATE,
  mood VARCHAR(50),
  notes TEXT,
  acidity TINYINT DEFAULT 0,
  bitterness TINYINT DEFAULT 0,
  sweetness TINYINT DEFAULT 0,
  body TINYINT DEFAULT 0,
  aroma TINYINT DEFAULT 0,
  aftertaste TINYINT DEFAULT 0,
  ai_summary TEXT,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME NULL
);
```

### 8.3 flavor_tags

```sql
CREATE TABLE flavor_tags (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) UNIQUE,
  label VARCHAR(100),
  color VARCHAR(50),
  created_at DATETIME,
  updated_at DATETIME
);
```

### 8.4 coffee_log_flavor_tags

```sql
CREATE TABLE coffee_log_flavor_tags (
  coffee_log_id BIGINT,
  flavor_tag_id BIGINT
);
```

### 8.5 card_templates

```sql
CREATE TABLE card_templates (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100),
  preview_url VARCHAR(500),
  layout_type VARCHAR(50),
  is_paid BOOLEAN DEFAULT FALSE,
  created_at DATETIME,
  updated_at DATETIME
);
```

---

## 9. API 返回格式

### 9.1 成功响应

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 9.2 失败响应

```json
{
  "code": 40001,
  "message": "error message",
  "data": null
}
```

### 9.3 分页响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 100
    }
  }
}
```

---

## 10. 关键接口设计

### 10.1 Auth

#### 注册

```http
POST /api/v1/auth/register
```

请求：

```json
{
  "email": "test@example.com",
  "password": "123456",
  "nickname": "Coffee Lover"
}
```

#### 登录

```http
POST /api/v1/auth/login
```

请求：

```json
{
  "email": "test@example.com",
  "password": "123456"
}
```

返回：

```json
{
  "token": "...",
  "user": {}
}
```

### 10.2 Coffee Logs

#### 创建记录

```http
POST /api/v1/coffee-logs
```

请求：

```json
{
  "coffee_name": "Ethiopia Yirgacheffe",
  "coffee_type": "Pour Over",
  "shop_name": "Blue Bottle",
  "location": "Shanghai",
  "image_url": "",
  "drink_date": "2026-05-21",
  "mood": "calm",
  "notes": "明亮、干净、像柑橘茶",
  "acidity": 4,
  "bitterness": 1,
  "sweetness": 3,
  "body": 2,
  "aroma": 5,
  "aftertaste": 4,
  "flavor_tag_ids": [1, 2, 4]
}
```

#### 查询列表

```http
GET /api/v1/coffee-logs?page=1&page_size=10&month=2026-05&coffee_type=Latte&tag_id=1
```

#### 查询详情

```http
GET /api/v1/coffee-logs/:id
```

#### 更新记录

```http
PUT /api/v1/coffee-logs/:id
```

#### 删除记录

```http
DELETE /api/v1/coffee-logs/:id
```

### 10.3 Stats

#### Overview

```http
GET /api/v1/stats/overview
```

返回：

```json
{
  "month_count": 18,
  "total_count": 120,
  "favorite_coffee_type": "Latte",
  "favorite_flavor_tag": "nutty"
}
```

#### Flavor Profile

```http
GET /api/v1/stats/flavor-profile
```

返回：

```json
{
  "acidity": 3.2,
  "bitterness": 1.8,
  "sweetness": 3.6,
  "body": 2.9,
  "aroma": 4.1,
  "aftertaste": 3.4
}
```

#### Monthly

```http
GET /api/v1/stats/monthly
```

返回：

```json
[
  {
    "month": "2026-01",
    "count": 12
  }
]
```

### 10.4 AI

```http
POST /api/v1/ai/flavor-summary
```

请求：

```json
{
  "coffee_type": "Pour Over",
  "tags": ["citrus", "floral"],
  "acidity": 4,
  "bitterness": 1,
  "sweetness": 3,
  "body": 2,
  "aroma": 5,
  "aftertaste": 4
}
```

返回：

```json
{
  "summary": "今天这杯咖啡带有明亮的柑橘酸质和轻盈花香，整体干净清爽，适合一个安静的下午。"
}
```

---

## 11. UI 设计方案

### 11.1 整体风格

整体采用 Nordic Minimal + Editorial。

也就是：

- 北欧极简
- 高级生活方式杂志风
- 温暖克制
- 大量留白
- 图片优先
- 排版驱动

### 11.2 配色

| 名称 | 色值 |
| --- | --- |
| Warm White | `#F7F3EC` |
| Cream | `#EFE7DA` |
| Latte Beige | `#D7C4A8` |
| Coffee Brown | `#7A5638` |
| Deep Espresso | `#2A1E17` |
| Charcoal | `#1E1E1E` |
| Soft Gray | `#A8A29A` |

### 11.3 布局原则

- 大量留白
- 大图优先
- 非对称排版
- 杂志式层级
- 避免普通 Dashboard
- 避免普通 SaaS 风
- 避免 Material Design 风
- 避免廉价渐变和重阴影

---

## 12. Docker Compose 方案

`docker-compose.yml` 需要包含：

- `mysql`：MySQL 8
- `redis`：Redis
- `backend`：Go Gin 服务
- `frontend`：Vue Vite 前端服务

建议端口：

- Frontend：`5173`
- Backend：`8080`
- MySQL：`3306`
- Redis：`6379`

Swagger 地址：

```text
http://localhost:8080/swagger/index.html
```

---

## 13. 环境变量

```env
APP_ENV=development
APP_PORT=8080

MYSQL_HOST=mysql
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=123456
MYSQL_DATABASE=my_coffee_log

REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=

JWT_SECRET=please_change_me
JWT_EXPIRE_HOURS=168

OPENAI_API_KEY=
OPENAI_BASE_URL=
OPENAI_MODEL=
```

---

## 14. 开发阶段规划

## 第一阶段：项目初始化与用户系统

### 目标

完成基础工程和登录闭环。

### 后端任务

- 初始化 Go Gin 项目
- 配置 Viper
- 配置 Zap 日志
- 配置 MySQL 连接
- 配置 Redis 连接
- 实现 CORS 中间件
- 实现 Recovery 中间件
- 实现日志中间件
- 实现 JWT 工具
- 实现 bcrypt 密码加密
- 实现用户注册
- 实现用户登录
- 实现获取当前用户信息
- 配置 Swagger

### 前端任务

- 初始化 Vue 3 + Vite + TypeScript
- 配置 Tailwind CSS
- 配置 Vue Router
- 配置 Pinia
- 实现 Login 页面
- 实现 Register 页面
- 实现路由守卫
- 实现 Auth Store
- 封装 API Client

### 验收结果

- 用户可以注册
- 用户可以登录
- 登录后可以进入首页
- JWT 鉴权可用
- Swagger 可访问
- Docker Compose 可以启动基础服务

---

## 第二阶段：Coffee Log CRUD 与 Timeline

### 目标

完成核心咖啡记录能力。

### 后端任务

- 实现 CoffeeLog 模型
- 实现 FlavorTag 模型
- 实现多对多关系
- 初始化默认风味标签
- 实现 Coffee Log 创建接口
- 实现 Coffee Log 列表接口
- 实现 Coffee Log 详情接口
- 实现 Coffee Log 更新接口
- 实现 Coffee Log 删除接口
- 实现分页与筛选
- 确保所有用户数据按 `user_id` 隔离

### 前端任务

- 实现 CreateCoffeeLog 页面
- 实现三步式记录流程
- 实现风味强度滑动条
- 实现风味标签选择
- 实现 Timeline 页面
- 实现月份、类型、标签筛选
- 实现 CoffeeDetail 基础页面

### 验收结果

- 登录后可以创建咖啡记录
- 可以查看自己的咖啡记录列表
- 可以查看记录详情
- 可以编辑和删除自己的记录
- 不允许访问其他用户记录

---

## 第三阶段：AI Mock、Stats 与雷达图

### 目标

让产品从普通记录变成咖啡生活手账。

### 后端任务

- 实现 `internal/service/ai_service.go`
- 实现 AI Mock 风味总结
- 实现 `POST /api/v1/ai/flavor-summary`
- 创建 Coffee Log 时自动生成 AI Summary
- 实现 `GET /api/v1/stats/overview`
- 实现 `GET /api/v1/stats/flavor-profile`
- 实现 `GET /api/v1/stats/monthly`

### 前端任务

- 实现 SVG 风味雷达图组件
- 在 CoffeeDetail 展示雷达图
- 在 CoffeeDetail 展示 AI 文案
- 实现 Stats 页面
- 优化 Home 页面数据展示

### 验收结果

- 创建记录后可生成 AI Mock 风味总结
- 咖啡详情页展示雷达图
- 统计页面展示风味偏好和月度统计

---

## 第四阶段：分享卡片与 UI 打磨

### 目标

完成产品传播功能和最终视觉质感。

### 后端任务

- 完善 Swagger 文档
- 完善基础测试或接口测试示例
- 完善 README
- 完善 Dockerfile 和 Docker Compose

### 前端任务

- 实现 ShareCardPreview
- 实现 CoffeeShareCard
- 支持 1:1、3:4、9:16 三种尺寸
- 使用 html2canvas 导出图片
- 支持下载图片到本地
- 打磨 Home 页面杂志感
- 打磨 Timeline 页面日记感
- 打磨 CoffeeDetail 页面高级感
- 打磨 Login / Register 品牌入口感

### 验收结果

- 可以生成分享卡片图片
- 整体 UI 具有高级生活方式杂志风
- 项目可通过 Docker Compose 启动
- README 完整
- Swagger 文档可访问

---

## 15. MVP 验收标准

MVP 必须满足：

1. 用户可以注册、登录。
2. 登录后可以创建咖啡记录。
3. 可以查看咖啡记录列表。
4. 可以查看咖啡详情。
5. 可以编辑、删除自己的咖啡记录。
6. 咖啡详情展示风味雷达图。
7. 可以生成 AI Mock 风味总结。
8. 可以生成分享卡片图片。
9. 可以查看简单统计。
10. 前后端可以通过 Docker Compose 启动。
11. Swagger 文档可访问。
12. README 完整。

---

## 16. 关键技术风险与建议

### 16.1 图片上传

当前需求字段是 `image_url`，但前端要求上传图片。

MVP 建议：

- 第一版支持图片 URL 或本地预览。
- 后续增加 `/api/v1/upload` 接口。
- 生产环境再接对象存储。

### 16.2 AI 接入

MVP 不建议直接接真实 OpenAI。

建议：

- 先完成 Mock AI Service。
- 保持接口抽象。
- 后续通过配置切换真实 AI Provider。

### 16.3 分享卡片导出

html2canvas 对跨域图片比较敏感。

建议：

- MVP 使用同源图片或允许用户输入可跨域访问的图片。
- 后续图片上传到后端或对象存储后统一处理 CORS。

### 16.4 数据迁移

MVP 可使用 GORM AutoMigrate 快速开发。

后续建议：

- 引入正式迁移工具。
- 将默认风味标签初始化逻辑独立成 seed。

### 16.5 UI 风格控制

最大风险是做成普通后台或 SaaS Dashboard。

建议：

- 少用传统卡片网格。
- 少用重阴影。
- 少用炫彩渐变。
- 强化图片、文字、留白和非对称排版。

---

## 17. README 需要包含的内容

最终 README 必须包含：

- 项目介绍
- 技术栈
- 目录结构
- 环境变量
- Docker Compose 启动方式
- 后端启动方式
- 前端启动方式
- 数据库迁移 / 初始化说明
- Swagger 文档地址
- 默认初始化风味标签
- 后续开发计划

---

## 18. 推荐开发顺序总结

### 阶段一

- 初始化 Go Gin 项目
- 初始化 Vue 3 项目
- 完成 Docker Compose
- 完成 MySQL / Redis
- 完成用户注册登录

### 阶段二

- 完成 Coffee Log CRUD
- 完成 Flavor Tag 初始化
- 完成 Timeline

### 阶段三

- 完成 Stats
- 完成 AI Mock Summary
- 完成雷达图

### 阶段四

- 完成分享卡片
- 优化 UI
- 完成 README 和 Swagger

---

## 19. 最终目标

最终交付的 My Coffee Log 应该是一款：

> 高级生活方式杂志风的咖啡记录 App。

它不应该只是普通 CRUD 工具。

它应该帮助用户把每天的一杯咖啡变成一种可回看、可分享、可沉淀的生活方式记录。
