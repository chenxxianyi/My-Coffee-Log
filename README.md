# ☕ My Coffee Log

> 面向日常咖啡爱好者的极简咖啡生活记录 App — Nordic Minimal + Editorial 风格的生活方式咖啡手账

---

## ✨ 功能概览

| 模块 | 功能 |
|------|------|
| 🔐 用户系统 | 注册、登录、JWT 鉴权、个人资料编辑 |
| 📝 咖啡日志 | 创建/查看/删除冲煮记录，6 维风味参数滑块 |
| 🏷️ 风味标签 | 10 种预设风味标签（花香、柑橘、莓果…），后端种子数据 + API 驱动 |
| 🤖 AI 风味总结 | 基于感官参数自动生成 Editorial 风味文字摘要 |
| 📊 风味雷达图 | SVG 实时渲染 6 维风味指纹雷达图 |
| 📅 时间线 | 按月份分组浏览历史冲煮记录 |
| 📈 统计面板 | 总览、风味轮廓、月度冲煮频次 |
| 🃏 分享卡片 | 生成精美咖啡手账卡片，支持导出高清 PNG |
| 👤 个人中心 | 昵称修改、**本地相片上传头像**、数据导出、退出登录 |
| 📷 图片上传 | 本地相片/手机拍摄直传，自动设为日志封面或头像 |

---

## 🏗️ 技术栈

### 前端

| 技术 | 说明 |
|------|------|
| Vue 3 | Composition API + `<script setup>` |
| Vite 5 | 极速开发服务器与构建 |
| TypeScript | 全量类型安全 |
| Pinia | 状态管理（auth / coffeeLog） |
| Vue Router 4 | 路由 + 导航守卫 |
| Tailwind CSS 3 | 原子化 CSS，Nordic Minimal 主题 |
| Axios | HTTP 请求 + JWT 拦截器 |
| Lucide Icons | 精致线性图标 |
| html2canvas | 分享卡片 PNG 导出 |
| SVG Radar Chart | 自定义风味雷达图组件 |

### 后端

| 技术 | 说明 |
|------|------|
| Go 1.22+ | 高性能后端语言 |
| Gin | HTTP 框架 |
| GORM | MySQL ORM |
| MySQL 8 | 主数据库 |
| Redis 7 | 缓存（可选） |
| JWT (golang-jwt/v5) | 无状态认证 |
| Zap | 结构化日志 |
| Viper | 配置管理 |
| Validator | 请求参数校验 |

### 基础设施

| 技术 | 说明 |
|------|------|
| Docker Compose | 一键启动 MySQL + Redis + 后端 + 前端 |
| Nginx | 生产环境反向代理 |

---

## 📁 项目结构

```
My Coffee Log/
├── frontend/                    # 前端项目
│   ├── src/
│   │   ├── api/                 # API 请求模块
│   │   │   ├── request.ts       # Axios 封装 + JWT 拦截器
│   │   │   ├── auth.ts          # 认证接口
│   │   │   ├── coffeeLog.ts     # 咖啡日志接口
│   │   │   ├── stats.ts         # 统计接口
│   │   │   └── flavorTag.ts     # 风味标签接口
│   │   ├── components/          # 公共组件
│   │   │   └── charts/          # 风味雷达图
│   │   ├── pages/               # 页面组件
│   │   │   ├── Splash.vue       # 启动页
│   │   │   ├── Login.vue        # 登录
│   │   │   ├── Register.vue     # 注册
│   │   │   ├── Home.vue         # 首页
│   │   │   ├── CreateCoffeeLog.vue  # 创建日志
│   │   │   ├── CoffeeDetail.vue     # 日志详情
│   │   │   ├── Timeline.vue     # 时间线
│   │   │   ├── Stats.vue        # 统计
│   │   │   ├── ShareCard.vue    # 分享卡片
│   │   │   └── Profile.vue     # 个人中心
│   │   ├── router/              # 路由配置 + 守卫
│   │   ├── stores/              # Pinia 状态仓库
│   │   │   ├── auth.ts          # 认证状态
│   │   │   └── coffeeLog.ts     # 日志状态
│   │   └── styles/              # 全局样式
│   ├── tailwind.config.js
│   ├── vite.config.ts           # Vite 配置 + API 代理
│   └── package.json
│
├── backend/                     # 后端项目
│   ├── cmd/server/              # 入口 main.go
│   ├── internal/
│   │   ├── config/              # Viper 配置加载
│   │   ├── database/            # MySQL / Redis 初始化
│   │   ├── handler/             # HTTP 处理器
│   │   │   ├── auth_handler.go
│   │   │   ├── user_handler.go
│   │   │   ├── coffee_log_handler.go
│   │   │   ├── stats_handler.go
│   │   │   ├── ai_handler.go
│   │   │   ├── flavor_tag_handler.go
│   │   │   └── upload_handler.go
│   │   ├── service/             # 业务逻辑层
│   │   ├── repository/          # 数据访问层
│   │   ├── model/               # 数据模型
│   │   ├── middleware/          # 中间件（Auth / CORS / Logger）
│   │   ├── response/            # 统一响应格式
│   │   ├── router/              # 路由注册
│   │   └── utils/               # 工具函数
│   ├── uploads/                 # 用户上传文件目录
│   ├── .env                     # 本地环境变量
│   ├── .env.example             # 环境变量示例
│   └── go.mod
│
├── docker-compose.yml           # 容器编排
└── 设计文件/                     # UI 设计稿
```

---

## 🚀 快速开始

### 前置条件

- Go 1.22+
- Node.js 18+
- MySQL 8.0
- Redis 7（可选，缺失时后端仍可正常运行）

### 方式一：本地开发

**1. 启动 MySQL**

确保 MySQL 服务运行中，创建数据库：

```sql
CREATE DATABASE my_coffee_log CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

**2. 配置后端环境变量**

```bash
cd backend
cp .env.example .env
# 编辑 .env，填入你的 MySQL 连接信息
```

`.env` 示例：

```env
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=123456
MYSQL_DBNAME=my_coffee_log
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
JWT_SECRET=your-jwt-secret-key
APP_PORT=8080
```

**3. 启动后端**

```bash
cd backend
go run ./cmd/server
```

后端启动后：
- API 地址：`http://localhost:8080/api/v1`
- 上传文件访问：`http://localhost:8080/uploads/`
- 首次启动会自动迁移表结构并种子风味标签

**4. 启动前端**

```bash
cd frontend
npm install
npm run dev
```

前端启动后访问：`http://localhost:5199`

> Vite 开发服务器已配置代理，`/api` 和 `/uploads` 请求自动转发至后端 `:8080`。

### 方式二：Docker Compose 一键启动

```bash
docker-compose up --build
```

启动后：
- 前端：`http://localhost:5173`
- 后端 API：`http://localhost:8080/api/v1`

---

## 📡 API 端点

### 公开接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/auth/register` | 用户注册 |
| POST | `/api/v1/auth/login` | 用户登录 |

### 认证接口（需 Bearer Token）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/users/me` | 获取当前用户信息 |
| PUT | `/api/v1/users/me` | 更新当前用户信息 |
| POST | `/api/v1/coffee-logs` | 创建咖啡日志 |
| GET | `/api/v1/coffee-logs` | 获取日志列表（分页） |
| GET | `/api/v1/coffee-logs/:id` | 获取日志详情 |
| PUT | `/api/v1/coffee-logs/:id` | 更新日志 |
| DELETE | `/api/v1/coffee-logs/:id` | 删除日志 |
| GET | `/api/v1/stats/overview` | 统计总览 |
| GET | `/api/v1/stats/flavor-profile` | 风味轮廓统计 |
| GET | `/api/v1/stats/monthly` | 月度冲煮统计 |
| POST | `/api/v1/ai/flavor-summary` | AI 风味总结生成 |
| GET | `/api/v1/flavor-tags` | 获取所有风味标签 |
| POST | `/api/v1/upload` | 上传图片文件 |

### 统一响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": { ... }
}
```

---

## 🔒 认证机制

- 登录/注册成功后后端返回 JWT Token
- 前端将 Token 存储于 `localStorage`
- 所有认证接口请求头携带：`Authorization: Bearer <token>`
- Token 过期或无效时，前端自动跳转至登录页
- 所有用户数据按 `user_id` 严格隔离

### 测试账号

| 字段 | 值 |
|------|-----|
| 邮箱 | `admin@mycoffeelog.com` |
| 密码 | `admin123` |
| 昵称 | Admin |

---

## 🎨 设计风格

**Nordic Minimal + Editorial** — 北欧极简杂志风

- 温暖米白底色 `#F7F3EC`
- 浓缩咖文字 `#2A1E17`
- 衬线字体标题 + 无衬线正文
- 微量双线边框（double-border）
- 大量留白 + 紧凑信息密度交替
- 风味雷达图采用低饱和度配色

---

## 📄 License

MIT
#   M y - C o f f e e - L o g  
 