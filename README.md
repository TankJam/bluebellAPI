# bluebellAPI
bluebellAPI 服务

# 服务端架构
Gin + Gorm + MySQL + Redis + Vue + Docker 

# GIT管理
- 项目中长期存在的两个分支
    - master：主分支，负责记录上线版本的迭代，该分支代码与线上代码是完全一致的。
    - develop：开发分支，该分支记录相对稳定的版本，所有的feature分支和bugfix分支都从该分支创建。

- 其它分支为短期分支，其完成功能开发之后需要删除
    - feature/*：特性（功能）分支，用于开发新的功能，不同的功能创建不同的功能分支，功能分支开发完成并自测通过之后，需要合并到 develop 分支，之后删除该分支。
    - bugfix/*：bug修复分支，用于修复不紧急的bug，普通bug均需要创建bugfix分支开发，开发完成自测没问题后合并到 develop 分支后，删除该分支。
    - release/*：发布分支，用于代码上线准备，该分支从develop分支创建，创建之后由测试同学发布到测试环境进行测试，测试过程中发现bug需要开发人员在该release分支上进行bug修复，所有bug修复完后，在上线之前，需要合并该release分支到master分支和develop分支。
    - hotfix/*：紧急bug修复分支，该分支只有在紧急情况下使用，从master分支创建，用于紧急修复线上bug，修复完成后，需要合并该分支到master分支以便上线，同时需要再合并到develop分支。
  
