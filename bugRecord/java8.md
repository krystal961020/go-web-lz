# Java8使用过程中碰到的问题

### for循环问题

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

### java代码大量if-else

**优化前**

```java
    private void buildOddsInfo(BaseSportConfig baseSportConfig, JSONObject jsonObject) {
        if (null != jsonObject.getJSONArray(ALL_WIN.getCode())) {
            buildWinOddsInfo(baseSportConfig, jsonObject, ALL_WIN);
        } else if (null != jsonObject.getJSONArray(ALL_BALL.getCode())) {
            buildOtherOddsInfo(baseSportConfig, jsonObject, ALL_BALL, HOME_BALL, AWAY_BALL);
        } else if (null != jsonObject.getJSONArray(ALL_BS.getCode())) {
            buildOtherOddsInfo(baseSportConfig, jsonObject, ALL_BS, BIG, SMALL);
        } else if (null != jsonObject.getJSONArray(HALF_WIN.getCode())) {
            buildWinOddsInfo(baseSportConfig, jsonObject, HALF_WIN);
        } else if (null != jsonObject.getJSONArray(HALF_BALL.getCode())) {
            buildOtherOddsInfo(baseSportConfig, jsonObject, HALF_BALL, HOME_BALL, AWAY_BALL);
        } else if (null != jsonObject.getJSONArray(HALF_BS.getCode())) {
            buildOtherOddsInfo(baseSportConfig, jsonObject, HALF_BS, BIG, SMALL);
        } else {
            baseSportConfig.setOddsType(HORIZONTAL_LINE);
            baseSportConfig.setOddsInfo(new HashMap<>(1));
        }
    }
```

**优化后**
```java
感觉好像没办法,但是看着好垃圾，哈哈哈哈
```