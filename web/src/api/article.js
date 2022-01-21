import { Post } from "@/service/index.js";
import { ElMessage } from "element-plus";
import { JudgeRequestStatus } from "@/utils";
import { UpLoadOS } from "@/api/cos.js";
export async function uploadArticle(article, config) {
  try {
    const { articleStorageType } = article;
    const upfunc = articleStorageType === "oos" ? uploadOos : uploadService;
    const upResult = await upfunc(article, config);
    return upResult;
  } catch (e) {
    console.error("上传文章错误", e);
  }
}

export function uploadOos(article, { oos }) {
  return new Promise((resolve, reject) => {
    const key = oos.slice(1).join("/");
    UpLoadOS(
      {
        Bucket: oos[0],
        Region: "ap-nanjing",
        Key: `${key}/${article.fileName}`,
        Body: article.articleContext,
      },
      function (err, data) {
        if (data) {
          if (data.statusCode === 200) {
            createArticle({
              ...article,
              articleUrl: data.Location,
            })
              .then(function (article) {
                console.log(1, article);
                resolve(true);
              })
              .catch((error) => {
                reject(false);
              });
          }
        } else {
          reject(err);
        }
      }
    );
  });
}

async function uploadService(article) {
  return new Promise((resolve, reject) => {
    resolve("文件上传");
  });
}

function createArticle(article) {
  return new Promise((resolve, reject) => {
    Post("/article/create_article", article)
      .then((response) => {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(response.data);
          },
          (code, msg) => {
            ElMessage.error(msg);
            reject({});
          }
        );
      })
      .catch((err) => {
        reject(err);
      });
  });
}
