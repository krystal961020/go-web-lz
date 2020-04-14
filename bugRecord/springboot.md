## SpringBoot

**@RestControllerAdvice的优先级要比在代码中的try catch的优先级要高**

```java
@Slf4j
@RestControllerAdvice
public class ExceptionAdviceCore extends ResponseEntityExceptionHandler {

    @ExceptionHandler(ServiceException.class)
    public ResponseVO<T> serviceExceptionAdvice(ServiceException serviceException) {
        log.error("controller serviceException error ", serviceException);
        return ResponseHelper.error(serviceException.getMsg());
    }
}
```

```java
@PostMapping("/v1/queryOperationActivityInfoList")
public ResponseVO queryOperationActivityInfoList(@RequestBody OperationActivityInfoDto operationActivityInfo, HttpServletRequest request) {
        try {
            return ResponseHelper.success(operationActivityInfoService.queryOperationActivityInfoList(operationActivityInfo, request));
        } catch (ServiceException e) {
            log.error("分页条件查询运营-活动管理异常！ex={}", e.getMessage(), e);
            return ResponseHelper.error(e.getMessage());
        }
}
```

以上代码是不会执行到

```java
log.error("分页条件查询运营-活动管理异常！ex={}", e.getMessage(), e);
```

**springboot缓存 @Catchable与@CatchEvict**
```java
    @Override
    @CacheEvict(value = "operationSponsoredConfig",key = "#param")
    public int updateStatusById(SponsoredConfigUpdateStatusDto param) throws ServiceException {
        return operationSponsoredConfigMapper.updateStatusById(param);
    }

    /**
     * @Description 分页查询赞助配置对象
     * @Author:jing
     */
    @Override
    @Cacheable(value = "operationSponsoredConfig", key = "#param", cacheManager = "cacheManager")
    public PageInfo queryOperationSponsoredConfigList(SponsoredConfigPageDto param) throws ServiceException {
        PageHelper.startPage(param.getPageNum(), param.getPageSize());
        List<OperationSponsoredConfig> operationSponsoredConfigs = operationSponsoredConfigMapper.queryOperationSponsoredConfigList(param);
        OssPathUtil.convertPath(operationSponsoredConfigs, SHARE_URL, SPONSORED_ICON_URL, SPONSORED_URL);
        return new PageInfo(operationSponsoredConfigs);
    }

    @Override
    @CacheEvict(value = "operationSponsoredConfig",key = "#param")
    public void updateSponsoredByTime(HeaderDto param) throws ServiceException {
        operationSponsoredConfigMapper.updateSponsoredByTime();
    }
```
***如果@CacheEvict在统一个实现类中引用同一个value那么该缓存将不生效，应该统一放在controller层处理***

