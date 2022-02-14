import { Post, confirmStatus, Get } from "@/utils/service";
import { article, state, articleList } from "@/types/response";
import {
  articleStorageType,
  CosTempKeyRequest,
  article_req,
} from "@/types/request";
import { putObject } from "./cos";
export async function getArticleList(
  limit?: number,
  offset?: number
): Promise<articleList> {
  return new Promise<articleList>((resolve, reject) => {
    Get(`article/get_article_list?limit${limit}&offset${offset}`)
      .then((article) => {
        const state = confirmStatus(<state>article.data.state);
        if (state) {
          resolve(article.data.data);
        } else {
          reject(article.data.state);
        }
      })
      .catch(reject);
  });
}

export async function createArticle(
  art: article_req,
  opt: string,
  key: Array<string>
): Promise<boolean> {
  const cosOpt: CosTempKeyRequest = {
    Region: "ap-nanjing",
    Bucket: "bloghtml-1301735126",
    action: "",
  };
  const path = key.concat(art.articleName).join("/");
  return new Promise<boolean>((resolve, reject) => {
    if (opt === articleStorageType.oos) {
      putObject(cosOpt, art.articleContext, path, (err, data) => {
        if (err) {
          reject(err);
        } else {
          if (data.statusCode === 200) {
            art.articleUrl = data.Location;
            postArticle(art).then(resolve).catch(reject);
          } else {
            reject(data);
          }
        }
      });
    } else {
      postArticle(art).then(resolve).catch(reject);
    }
  });
}

export async function postArticle(art: article_req): Promise<boolean> {
  return new Promise<boolean>((resolve, reject) => {
    Post("article/create_article", art)
      .then((res) => {
        const state = confirmStatus(<state>res.data.state);
        if (state) {
          resolve(res.data.data);
        } else {
          reject(res.data.state);
        }
      })
      .catch(reject);
  });
}
