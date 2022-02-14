import { getCosTemKey } from "@/api/cos";
import { article, state, cosTemKey } from "@/types/response";
import { CosTempKeyRequest } from "@/types/request";
import COS from "cos-js-sdk-v5";

export async function createCosInstance(opt: CosTempKeyRequest) :Promise<any> {
  try {
    const key = await getCosTemKey(opt);
    return new COS({
      getAuthorization: function (options, callback) {
        callback({
          TmpSecretId: key.Credentials.TmpSecretId,
          TmpSecretKey: key.Credentials.TmpSecretKey,
          SecurityToken: key.Credentials.Token,
          StartTime: key.StartTime, // 时间戳，单位秒，如：1580000000
          ExpiredTime: key.ExpiredTime, // 时间戳，单位秒，如：1580000000
        });
      },
    });
  } catch (err) {
    return err
  }
}
