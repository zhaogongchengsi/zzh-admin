const { contextBridge } = require('electron')
// 给当前执行器上下文 挂在api
contextBridge.exposeInMainWorld(
    'electron',
    {
      doThing: () => console.log(1),
      myPromises: [Promise.resolve(), Promise.reject(new Error('whoops'))],
      anAsyncFunction: async () => 123,
      data: {
        myFlags: ['a', 'b', 'c'],
        bootTime: 1234
      },
      nestedAPI: {
        evenDeeper: {
          youCanDoThisAsMuchAsYouWant: {
            fn: () => ({
              returnData: 123
            })
          }
        }
      }
    }
)