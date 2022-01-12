import { createCosInstance } from "@/service/os.js";

// Bucket: "examplebucket-1250000000" /* 填入您自己的存储桶，必须字段 */,
// Region: "COS_REGION" /* 存储桶所在地域，例如ap-beijing，必须字段 */,
// Key: "1.jpg" /* 存储在桶里的对象键（例如1.jpg，a/b/test.txt），必须字段 */,
// StorageClass: "STANDARD",
// Body: fileObject, // 上传文件对象
export async function UpLoadOS(data, callback, load = true) {
  let cosinstance = await createCosInstance(data);
  cosinstance.putObject(data, callback);
}
