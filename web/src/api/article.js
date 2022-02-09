import { Post, Get } from "@/service/index.js";
import { ElMessage } from "element-plus";
import { JudgeRequestStatus } from "@/utils";
import { UpLoadOS } from "@/api/cos.js";
export async function uploadArticle(article, config) {
  try {
    const { articleStorageType } = article;
    if (articleStorageType === "oos") {
      return await uploadOos(article, config);
    }
    return await createArticle(article);
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
          console.log(data);
          if (data.statusCode === 200) {
            article.articleContext = "";
            createArticle({
              ...article,
              articleUrl: data.Location,
            })
              .then(function (article) {
                resolve(article);
              })
              .catch((error) => {
                reject(error);
              });
          }
        } else {
          reject(err);
        }
      }
    );
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

export function getArticles(limit, offset) {
  return new Promise((resolve, reject) => {
    Get(`article/get_article_list?limit=${limit}&offset=${offset}`)
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
