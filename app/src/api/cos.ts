import { Post, confirmStatus, Get } from "@/utils/service";
import { article, state, cosTemKey } from "@/types/response";
import { CosTempKeyRequest } from "@/types/request";
import { createCosInstance } from "@/utils/cos";
export async function getCosTemKey(opt: CosTempKeyRequest): Promise<cosTemKey> {
  return new Promise<cosTemKey>((resolve, reject) => {
    Post("cos/get_temporary_key", opt)
      .then((res) => {
        const state = confirmStatus(<state>res.data.state);
        if (state) {
          resolve(res.data.data);
        } else {
          reject(res.data.state);
        }
      })
      .catch((err) => {
        console.error(err);
      });
  });
}

type putfn = (err: any, data: any) => void;
export async function putObject(
  opt: CosTempKeyRequest,
  obj: string | File,
  key: string,
  fn: putfn
) {
  const cos = await createCosInstance(opt);
  console.log('创建cos 实例', cos);
  cos.putObject(
    {
      Bucket: opt.Bucket /* 填入您自己的存储桶，必须字段 */,
      Region: opt.Region /* 存储桶所在地域，例如ap-beijing，必须字段 */,
      Key: key /* 存储在桶里的对象键（例如1.jpg，a/b/test.txt），必须字段 */,
      Body: obj, // 上传文件对象
      onProgress: function (progressData) {
        console.log(JSON.stringify(progressData));
      },
    },
    fn
  );
}
