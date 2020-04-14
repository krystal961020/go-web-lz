# git坑死我了

```git
BUG: There are unmerged index entries:
BUG: 2 kkcloud-site-admin-biz/src/main/java/com/kkcloud/site/admin/validate/DictUtils.java
BUG: merge-recursive.c:429: unmerged index entries in merge-recursive.c
```

```
git merge -s resolve branch_A(对应的分支)
```

**问题描述：由于本地merge代码后没有打出resolve标记所以git会存在对应缓存需要手动解决掉**