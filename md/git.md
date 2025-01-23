# Git 常用命令指南

## 1. Git 配置

### 1.1 设置用户信息
```bash
# 设置全局用户名
git config --global user.name "你的名字"

# 设置全局邮箱
git config --global user.email "你的邮箱@example.com"

# 查看全局配置
git config --list
```

### 1.2 生成SSH密钥
```bash
# 生成SSH密钥
ssh-keygen -t rsa -C "你的邮箱@example.com"

# 查看公钥内容
cat ~/.ssh/id_rsa.pub
```

## 2. 基本操作

### 2.1 仓库初始化和克隆
```bash
# 初始化新仓库
git init

# 克隆远程仓库
git clone <仓库地址>

# 克隆指定分支
git clone -b <分支名> <仓库地址>
```

### 2.2 添加和提交
```bash
# 查看仓库状态
git status

# 添加指定文件到暂存区
git add <文件名>

# 添加所有改动到暂存区
git add .

# 提交更改
git commit -m "提交说明"

# 修改最后一次提交信息
git commit --amend
```

### 2.3 分支操作
```bash
# 查看所有分支
git branch

# 创建新分支
git branch <分支名>

# 切换分支
git checkout <分支名>

# 创建并切换到新分支
git checkout -b <分支名>

# 删除分支
git branch -d <分支名>

# 强制删除分支
git branch -D <分支名>
```

## 3. 远程操作

### 3.1 远程仓库
```bash
# 查看远程仓库
git remote -v

# 添加远程仓库
git remote add origin <仓库地址>

# 推送到远程仓库
git push origin <分支名>

# 拉取远程更新
git pull origin <分支名>

# 获取远程更新但不合并
git fetch origin
```

### 3.2 标签管理
```bash
# 创建标签
git tag <标签名>

# 创建带注释的标签
git tag -a <标签名> -m "标签说明"

# 查看所有标签
git tag

# 推送标签到远程
git push origin <标签名>

# 推送所有标签
git push origin --tags
```

## 4. 撤销和回滚

### 4.1 撤销修改
```bash
# 撤销工作区修改
git checkout -- <文件名>

# 撤销暂存区修改
git reset HEAD <文件名>

# 回退到指定版本
git reset --hard <commit_id>
```

### 4.2 储藏和恢复
```bash
# 储藏当前修改
git stash

# 查看储藏列表
git stash list

# 恢复最近的储藏
git stash pop

# 删除储藏
git stash drop
```

## 5. 高级操作

### 5.1 合并和变基
```bash
# 合并分支
git merge <分支名>

# 变基操作
git rebase <分支名>

# 中断变基
git rebase --abort
```

### 5.2 日志查看
```bash
# 查看提交历史
git log

# 查看简洁历史
git log --oneline

# 查看分支图
git log --graph --oneline --all
```

### 5.3 差异比较
```bash
# 查看工作区和暂存区的差异
git diff

# 查看暂存区和最新提交的差异
git diff --cached

# 查看两个分支的差异
git diff <分支1> <分支2>
```

## 6. 常见问题解决

### 6.1 解决冲突
```bash
# 手动解决冲突后，添加文件
git add <冲突文件>

# 继续合并
git merge --continue

# 继续变基
git rebase --continue
```

### 6.2 清理仓库
```bash
# 清理未跟踪的文件
git clean -f

# 清理未跟踪的文件和目录
git clean -fd

# 清理忽略的文件
git clean -fX
```

## 7. 最佳实践

1. 经常性地提交代码，保持提交粒度适中
2. 提交信息要清晰明了，说明本次提交的主要改动
3. 使用 .gitignore 忽略不需要版本控制的文件
4. 定期从远程仓库拉取更新
5. 重要功能开发使用新分支
6. 及时处理冲突，保持代码库整洁

## 8. 安全建议

1. 不要将敏感信息（密码、密钥等）提交到代码库
2. 重要操作前先备份
3. 定期备份代码库
4. 谨慎使用 force push
5. 保护好你的 SSH 密钥
