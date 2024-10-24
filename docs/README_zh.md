<h1 align="center">ZBook</h1>
<p align="center">
  <a href="https://github.com/zizdlp/zbook-docs">文档</a> - <a href="https://github.com/zizdlp/zbook-helm-chart">Helm Chart</a>  - <a href="https://discord.com/channels/1250069935594536960/1250069935594536963">Discord</a> - <a href="https://www.youtube.com/channel/UC9D6VAJRoG7bD38dz8F9CSg">Youtube</a>
</p>

<div align="center">

[![Actions Status](https://github.com/zizdlp/zbook/workflows/BUILD_MAIN/badge.svg?branch=main)](https://github.com/zizdlp/zbook/actions)
[![Actions Status](https://github.com/zizdlp/zbook/workflows/TEST_BACKEND/badge.svg?branch=main)](https://github.com/zizdlp/zbook/actions)
[![Actions Status](https://github.com/zizdlp/zbook/workflows/TEST_FRONTEND/badge.svg?branch=main)](https://github.com/zizdlp/zbook/actions)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/zizdlp/zbook/tree/release.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/zizdlp/zbook/tree/release)

</div>

[English Version](README.md)

------

<p align="center">欢迎来到 ZBook，一个为团队设计的完全开源的全栈知识库管理软件。</p>
<p align="center">这个仓库包含用于渲染和服务 ZBook 的开源代码。</p>

<p align="center">
  <img alt="group_demo" src="./zbook_frontend/public/group_demo.png">
</p>

## 目录

- [入门指南](#getting-started)
- [功能](#features)
- [部署](#deployment)
- [许可](#license)
- [致谢](#acknowledgements)

## 入门指南

要运行这个项目的本地版本，请按照以下简单步骤操作。

### 先决条件

- Docker 和 Docker Compose

### 设置

1. 将仓库克隆到一个 公开 的 GitHub 仓库。如果你计划分发代码，请保持源码公开以符合 GNU GPLv3。

    ```shell
    git clone https://github.com/zizdlp/zbook.git
    ```

2. 构建并运行本地开发服务器

    ```shell
    make compose_pull
    ```

3. 然后在你的网页浏览器中打开 <http://localhost:3000/>

4. 🍻 要使用电子邮件服务和第三方账户登录，请先申请。详情请参阅 [ZBook 文档](https://github.com/zizdlp/zbook-docs).

### CI 和测试

所有的拉取请求都会进行视觉和性能测试，以防止回归问题。
ZBook 完全开源，并建立在 [Next.js](https://nextjs.org/)，[Golang](https://go.dev/) 之上.

### 贡献类型

我们鼓励你为 ZBook 做贡献，帮助我们构建最好的技术知识文档工具。如果你想快速做出贡献，请继续阅读以了解更多关于热门贡献的方法。

#### 翻译

ZBook 的用户界面使用一组翻译文件进行渲染，这些文件可以在 [`zbook_frontend/messages`](/zbook_frontend/messages/) 中找到。我们欢迎所有 UI 的额外翻译。

如果您想提交新的语言支持，请首先参考文档 [语言支持](https://github.com/zizdlp/zbook-docs/blob/main/Development/LanguageSupport.md)，跟随其中的相关步骤修改代码。此外，可以参考这些提交: `ec05f3a1d75d3f88619489a44f77104a37295ba3`,
`92e4b2933a9e23d08c15042b9b9085f4fea556f7`

#### 错误

遇到错误或找到你想修复的问题？帮助我们修复与 ZBook 相关的问题可以极大地改善所有人的体验。前往本仓库的问题部分，了解你可以帮助解决的错误类型。

## 功能

支持:

- **全栈**：利用 Next.js 和 Tailwind CSS 的前端，Golang gRPC 的后端服务，PostgreSQL 的数据库管理，MinIO 的存储，以及用于实时消息通知的 WebSocket。
- **多级权限管理**：支持多种存储库可见性选项，包括公开、仅登录、选定用户和仅创建者访问。
- **评论**
- **通知**
- **开源**

不支持:

ZBook 不支持**在线编辑**。我们认为 Git 在本地环境如 VS Code 和 Typora 中的工具足以用于编辑和协作内容。在线编辑不是必需的。此外，启用在线编辑需要授予 Git 仓库写权限，这可能带来**安全风险**。

### TODO

zbook计划支持如下功能：

- [x] 分支切换，动机是合并main分支之前，可以先查看特定分支的显示效果
- [x] mkdocs 风格的折叠式admonition
- [ ] 类似mkdocs的风格扩展
- [ ] github 风格的admonition

## 部署

由于隐私和其他原因（对于私人仓库，你可能需要输入访问令牌；虽然 GitHub 支持细粒度访问令牌，可以授予特定仓库的特定权限，例如只读），我们鼓励你自行部署 ZBook。你可以使用 Docker Compose 或 k8s 集群进行部署。详情请参阅 [ZBook 文档](https://github.com/zizdlp/zbook-docs)。

## 许可

根据[GNU GPLv3 License](https://github.com/zizdlp/zbook/LICENSE) 许可证 分发。

如果你计划分发代码，必须保持源码公开以符合 GNU GPLv3。
请参阅 LICENSE 了解更多信息。

## 致谢

没有以下项目，ZBook 将无法实现：

- [Next.js](https://nextjs.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [GoldMark](https://github.com/yuin/goldmark)

## 贡献者

<a href="https://github.com/zizdlp/zbook/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=zizdlp/zbook" />
</a>
