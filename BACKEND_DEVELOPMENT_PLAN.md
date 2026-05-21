# My Coffee Log 后端开发方案

## 1. 后端目标

My Coffee Log 后端负责提供用户认证、咖啡记录、风味标签、AI Mock 总结、数据统计和 Swagger 文档等能力。

后端必须采用清晰分层架构，保证业务逻辑、数据访问和 HTTP 处理职责分离。

核心原则：

- Handler 不直接操作数据库。
- Service 负责业务逻辑。
- Repository 负责数据访问。
- 所有用户数据必须按 `user_id` 隔离。
- 注册和登录以外的接口必须经过 JWT 鉴权。
- API 返回格式统一。

---

## 2. 后端技术栈

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

---

## 3. 后端目录结构

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

---

## 4. 分层职责

### 4.1 Handler 层

职责：

- 接收 HTTP 请求。
- 绑定和校验请求参数。
- 从上下文读取当前用户 ID。
- 调用 Service。
- 返回统一响应。

禁止：

- 直接访问数据库。
- 写复杂业务逻辑。

### 4.2 Service 层

职责：

- 处理业务逻辑。
- 编排多个 Repository。
- 处理密码校验、JWT 签发、AI 总结生成等业务动作。
- 保证用户数据隔离。

### 4.3 Repository 层

职责：

- 封装 GORM 查询。
- 实现创建、更新、删除、分页、筛选、统计等数据库操作。
- 不处理 HTTP 细节。

### 4.4 Model 层

职责：

- 定义 GORM 模型。
- 定义表字段。
- 定义模型关联关系。

### 4.5 Middleware 层

职责：

- JWT 鉴权。
- CORS。
- 请求日志。
- Panic Recovery。

### 4.6 Response 层

职责：

- 统一成功响应。
- 统一失败响应。
- 统一分页响应。

### 4.7 Utils 层

职责：

- JWT 工具。
- bcrypt 密码工具。
- 分页工具。
- Validator 参数校验工具。

---

## 5. 环境变量

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

## 6. 数据库设计

## 6.1 users

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

### GORM 关系

- User 与 CoffeeLog 是一对多关系。

---

## 6.2 coffee_logs

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

### 字段要求

风味维度范围：0-5。

- `acidity`
- `bitterness`
- `sweetness`
- `body`
- `aroma`
- `aftertaste`

### GORM 关系

- CoffeeLog 属于 User。
- CoffeeLog 与 FlavorTag 是多对多关系。

---

## 6.3 flavor_tags

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

### 默认风味标签

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

---

## 6.4 coffee_log_flavor_tags

```sql
CREATE TABLE coffee_log_flavor_tags (
  coffee_log_id BIGINT,
  flavor_tag_id BIGINT
);
```

---

## 6.5 card_templates

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

MVP 阶段可以先建表，不实现复杂付费模板逻辑。

---

## 7. GORM Model 要求

所有主模型需要包含：

- `ID`
- `CreatedAt`
- `UpdatedAt`
- `DeletedAt`

### User

字段：

- `ID`
- `Email`
- `Nickname`
- `PasswordHash`
- `AvatarURL`
- `CoffeeLogs`
- `CreatedAt`
- `UpdatedAt`
- `DeletedAt`

### CoffeeLog

字段：

- `ID`
- `UserID`
- `CoffeeName`
- `CoffeeType`
- `ShopName`
- `Location`
- `ImageURL`
- `DrinkDate`
- `Mood`
- `Notes`
- `Acidity`
- `Bitterness`
- `Sweetness`
- `Body`
- `Aroma`
- `Aftertaste`
- `AISummary`
- `FlavorTags`
- `CreatedAt`
- `UpdatedAt`
- `DeletedAt`

### FlavorTag

字段：

- `ID`
- `Name`
- `Label`
- `Color`
- `CreatedAt`
- `UpdatedAt`

### CardTemplate

字段：

- `ID`
- `Name`
- `PreviewURL`
- `LayoutType`
- `IsPaid`
- `CreatedAt`
- `UpdatedAt`

---

## 8. API 返回格式

### 8.1 成功响应

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 8.2 失败响应

```json
{
  "code": 40001,
  "message": "error message",
  "data": null
}
```

### 8.3 分页响应

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

## 9. 认证与用户模块

## 9.1 注册

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

处理流程：

1. 校验参数。
2. 校验邮箱是否已存在。
3. 使用 bcrypt 加密密码。
4. 创建用户。
5. 返回用户信息。

---

## 9.2 登录

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

处理流程：

1. 校验参数。
2. 根据邮箱查询用户。
3. 使用 bcrypt 校验密码。
4. 生成 JWT。
5. 返回 token 和用户信息。

---

## 9.3 获取当前用户

```http
GET /api/v1/users/me
```

要求：

- 需要 JWT 鉴权。
- 从上下文读取当前用户 ID。
- 返回当前用户信息。

---

## 9.4 更新当前用户

```http
PUT /api/v1/users/me
```

允许更新：

- `nickname`
- `avatar_url`

---

## 10. JWT 鉴权方案

### 10.1 Token 内容

JWT 建议包含：

- `user_id`
- `email`
- `exp`

### 10.2 Header 格式

```http
Authorization: Bearer <token>
```

### 10.3 中间件职责

- 读取 Authorization Header。
- 校验 Bearer Token。
- 解析 JWT。
- 将 `user_id` 写入 Gin Context。
- Token 无效时返回统一错误响应。

### 10.4 鉴权范围

不需要鉴权：

- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

需要鉴权：

- 其他全部业务接口。

---

## 11. Coffee Log 模块

## 11.1 创建记录

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

处理流程：

1. JWT 中间件获取当前用户 ID。
2. Handler 绑定请求参数。
3. Service 校验风味分值范围。
4. Service 查询并校验风味标签。
5. Service 调用 AIService 生成 Mock Summary。
6. Repository 创建 CoffeeLog。
7. Repository 绑定多对多 FlavorTags。
8. 返回完整记录。

---

## 11.2 查询列表

```http
GET /api/v1/coffee-logs?page=1&page_size=10&month=2026-05&coffee_type=Latte&tag_id=1
```

支持筛选：

- 分页
- 月份
- 咖啡类型
- 风味标签

默认排序：

```text
drink_date DESC, created_at DESC
```

权限要求：

- 必须只查询当前用户的记录。

---

## 11.3 查询详情

```http
GET /api/v1/coffee-logs/:id
```

要求：

- 需要 JWT。
- 必须校验 `user_id`。
- 只允许查看自己的记录。

---

## 11.4 更新记录

```http
PUT /api/v1/coffee-logs/:id
```

要求：

- 需要 JWT。
- 只允许更新自己的记录。
- 更新风味标签时需要同步多对多关系。
- 如果风味字段或标签变化，可以重新生成 AI Summary。

---

## 11.5 删除记录

```http
DELETE /api/v1/coffee-logs/:id
```

要求：

- 需要 JWT。
- 只允许删除自己的记录。
- 使用 GORM 软删除。

---

## 12. Flavor Tag 模块

### 职责

- 初始化默认风味标签。
- 支持 Coffee Log 绑定标签。
- 支持列表筛选和统计使用。

### 初始化策略

MVP 可在服务启动时执行 seed：

- 如果标签不存在，则创建。
- 如果标签已存在，则跳过。

后续可改为 migration/seed 命令。

---

## 13. AI 风味总结模块

## 13.1 接口

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

## 13.2 MVP 实现

后端本地 Mock 生成。

建议根据以下因素拼接文案：

- 咖啡类型
- 风味标签
- 酸度高低
- 甜感高低
- 苦感高低
- 香气高低
- 醇厚度
- 余韵

## 13.3 后续 OpenAI 预留

封装在：

```text
internal/service/ai_service.go
```

保留配置：

- `OPENAI_API_KEY`
- `OPENAI_BASE_URL`
- `OPENAI_MODEL`

后续可以通过配置决定使用 Mock Provider 还是真实 OpenAI Provider。

---

## 14. Stats 统计模块

## 14.1 Overview

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

统计逻辑：

- `month_count`：当前月份当前用户记录数。
- `total_count`：当前用户总记录数。
- `favorite_coffee_type`：当前用户出现次数最多的咖啡类型。
- `favorite_flavor_tag`：当前用户出现次数最多的风味标签。

---

## 14.2 Flavor Profile

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

统计逻辑：

- 计算当前用户所有 Coffee Log 的六个风味维度平均值。

---

## 14.3 Monthly

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

统计逻辑：

- 按月份聚合当前用户的 Coffee Log 数量。

---

## 15. 中间件方案

### 15.1 CORS

允许前端开发地址访问，例如：

- `http://localhost:5173`

需要支持：

- Authorization Header
- JSON 请求
- 常见 HTTP Method

### 15.2 Logger

使用 Zap 记录：

- 请求方法
- 请求路径
- 状态码
- 耗时
- 客户端 IP
- 错误信息

### 15.3 Recovery

捕获 panic，记录错误日志，并返回统一失败响应。

### 15.4 Auth

校验 JWT，写入当前用户上下文。

---

## 16. Docker Compose 方案

`docker-compose.yml` 需要包含：

- `mysql`：MySQL 8
- `redis`：Redis
- `backend`：Go Gin 服务
- `frontend`：Vue Vite 前端服务

建议端口：

- Backend：`8080`
- Frontend：`5173`
- MySQL：`3306`
- Redis：`6379`

Swagger 地址：

```text
http://localhost:8080/swagger/index.html
```

---

## 17. Swagger 文档

使用 Swaggo 生成接口文档。

要求：

- Auth 接口有请求和响应说明。
- Coffee Log 接口有请求和响应说明。
- Stats 接口有响应说明。
- AI 接口有请求和响应说明。
- 需要标注 JWT 鉴权。

访问地址：

```text
http://localhost:8080/swagger/index.html
```

---

## 18. 后端开发阶段

## 第一阶段：项目初始化与用户系统

### 任务

- 初始化 Go module。
- 初始化 Gin 服务。
- 配置 Viper。
- 配置 Zap。
- 配置 MySQL 连接。
- 配置 Redis 连接。
- 配置 GORM。
- 实现 User Model。
- 实现 Auth Repository。
- 实现 Auth Service。
- 实现 Auth Handler。
- 实现 User Handler。
- 实现 bcrypt 密码工具。
- 实现 JWT 工具。
- 实现 JWT 中间件。
- 实现 CORS、Logger、Recovery 中间件。
- 配置 Swagger。
- 编写 Dockerfile。
- 编写 `.env.example`。

### 验收

- 后端服务可以启动。
- 可以连接 MySQL 和 Redis。
- 用户可以注册。
- 用户可以登录。
- 登录后可以获取当前用户。
- Swagger 可以访问。

---

## 第二阶段：Coffee Log CRUD 与 Flavor Tag

### 任务

- 实现 CoffeeLog Model。
- 实现 FlavorTag Model。
- 实现 CardTemplate Model。
- 实现 CoffeeLog 与 FlavorTag 多对多关系。
- 实现默认风味标签初始化。
- 实现 CoffeeLog Repository。
- 实现 FlavorTag Repository。
- 实现 CoffeeLog Service。
- 实现 CoffeeLog Handler。
- 实现创建、列表、详情、更新、删除接口。
- 实现分页。
- 实现月份、咖啡类型、标签筛选。
- 确保所有 Coffee Log 操作按 `user_id` 隔离。

### 验收

- 登录用户可以创建咖啡记录。
- 可以分页查询自己的记录。
- 可以按月份、咖啡类型、标签筛选。
- 可以查看自己的记录详情。
- 可以编辑和删除自己的记录。
- 不能访问、修改或删除其他用户的记录。

---

## 第三阶段：AI Mock 与 Stats

### 任务

- 实现 AI Service。
- 实现 Mock 风味总结生成。
- 实现 AI Handler。
- 实现 `POST /api/v1/ai/flavor-summary`。
- Coffee Log 创建时自动生成 AI Summary。
- 实现 Stats Repository。
- 实现 Stats Service。
- 实现 Stats Handler。
- 实现 Overview 统计。
- 实现 Flavor Profile 统计。
- 实现 Monthly 统计。

### 验收

- 创建 Coffee Log 后自动生成 AI Summary。
- 可以单独调用 AI Summary 接口。
- 可以获取本月记录数、总记录数、最常喝类型、最常见标签。
- 可以获取风味平均值。
- 可以获取月度统计。

---

## 第四阶段：文档、测试和部署完善

### 任务

- 完善 Swagger 注释。
- 完善 Dockerfile。
- 完善 Docker Compose。
- 完善 README。
- 添加基础单元测试或接口测试示例。
- 补充数据库初始化说明。
- 补充默认风味标签说明。

### 验收

- `docker-compose up` 可启动前后端、MySQL、Redis。
- Swagger 文档完整可访问。
- README 包含启动说明。
- 有基础测试或接口测试示例。

---

## 19. 后端测试建议

### 19.1 单元测试

优先覆盖：

- 密码加密与校验。
- JWT 生成与解析。
- AI Mock Summary 生成。
- Service 层核心逻辑。

### 19.2 接口测试

建议提供示例：

- 注册
- 登录
- 获取当前用户
- 创建 Coffee Log
- 查询 Coffee Log 列表
- 查询 Coffee Log 详情
- 更新 Coffee Log
- 删除 Coffee Log
- 获取 Stats

---

## 20. 后端 MVP 验收标准

1. 用户可以注册、登录。
2. JWT 鉴权可用。
3. 登录用户可以获取当前用户信息。
4. 登录用户可以创建咖啡记录。
5. 登录用户可以查看自己的咖啡记录列表。
6. 登录用户可以查看自己的咖啡记录详情。
7. 登录用户可以编辑、删除自己的咖啡记录。
8. 所有用户数据按 `user_id` 隔离。
9. 创建记录后可以生成 AI Mock 风味总结。
10. Stats 接口可以返回基础统计。
11. Swagger 文档可访问。
12. 后端可以通过 Docker Compose 启动。

---

## 21. 关键技术风险

### 21.1 用户数据越权

所有 Coffee Log、Stats、Timeline 相关查询必须带 `user_id` 条件。

重点防护：

- 查询详情
- 更新记录
- 删除记录
- 统计数据

### 21.2 AI 接入复杂化

MVP 不接真实 OpenAI，只做 Mock。

AI Service 需要预留接口，避免后续重构过大。

### 21.3 数据迁移策略

MVP 可使用 GORM AutoMigrate。

后续建议引入正式 migration 工具，并将默认风味标签初始化独立成 seed。

### 21.4 图片上传

当前后端只需要支持 `image_url` 字段。

MVP 后续可以增加上传接口：

```http
POST /api/v1/upload
```

生产环境建议接对象存储。

---

## 22. 后端最终目标

后端最终需要成为一个稳定、清晰、可扩展的 Coffee Log API 服务。

它要支撑：

- 快速记录
- 用户数据隔离
- 风味总结
- 时间线查询
- 数据统计
- 分享卡片数据展示

核心要求不是功能堆叠，而是保证业务边界清晰、接口稳定、安全可靠。
