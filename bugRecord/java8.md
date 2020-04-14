# Java8使用过程中碰到的问题

```java
    /**
     * 判断基础联系信息json字符是否符合要求<br/>
     * 1、启用总数不能大于三个<br/>
     * 2、联系方式描述不能为空<br/>
     *
     * @param param 前端传入参数
     */
    private void checkContactBaseInfoJson(SiteBaseConfigDto param) {
        int countAble = 0;
        List<SiteBaseConfigInfo> baseInfo = param.getBaseInfo();
        for (SiteBaseConfigInfo siteBaseConfigInfo : baseInfo) {
           if(Objects.equals(siteBaseConfigInfo.getIsEnable(),ABLE.getCode()))) {
                Assert.isFalse(isBlank(siteBaseConfigInfo.getDeviceNum()), NUMBER_NOT_BE_NULL);
                countAble++;
            }
        }
        Assert.isFalse(countAble > MAX_ENABLE_COUNT, BASE_INFO_RATHER_THREE);
        param.setBaseInfos(JSON.toJSONString(baseInfo));
    }
```

```java
    /**
     * 判断基础联系信息json字符是否符合要求<br/>
     * 1、启用总数不能大于三个<br/>
     * 2、联系方式描述不能为空<br/>
     *
     * @param param 前端传入参数
     */
    private void checkContactBaseInfoJson(SiteBaseConfigDto param) {
        AtomicInteger countAble = new AtomicInteger();
        List<SiteBaseConfigInfo> baseInfo = param.getBaseInfo();
        baseInfo.stream().filter(siteBaseConfigInfo -> Objects.equals(siteBaseConfigInfo.getIsEnable(),ABLE.getCode())).forEach(siteBaseConfigInfo -> {
            Assert.isFalse(isBlank(siteBaseConfigInfo.getDeviceNum()), NUMBER_NOT_BE_NULL);
            countAble.getAndIncrement();
        });
        Assert.isFalse(countAble.get() > MAX_ENABLE_COUNT, BASE_INFO_RATHER_THREE);
        param.setBaseInfos(JSON.toJSONString(baseInfo));
    }
```

**以上第二段代码是不会进入foreach当中**