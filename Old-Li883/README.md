### 操作说明：

- 首先 fork [此仓库](https://github.com/tofar/bingyan-summer-camp2018)

- 在你的仓库操作的时候请不要对他人目录进行任何操作

- 你的操作权限仅限于你的目录，目录名字为你的 githubID，若仓库中没有你的目录请自行创建

- 提交 PR 的时候自行查看是否存在代码冲突，如果存在自行解决之后再提交 PR

- 提交 PR 是提交到 dev 分支，不是 master 分支

- 提交之后最好跟导师说一声，让导师及时检查

- 目录结构推荐如下：

  README.md   必须注明相关开发环境，以及一些特殊说明

  .gitignore  忽略一些不需要的文件

  client

  前端代码，如果不用webpack之类的打包的话，直接一个 dist 文件夹即可

  - dist

    编译打包后的文件目录

    - index.html
    - css
    - js

  - src

    源码

  server

  服务端代码

  - src

    源码