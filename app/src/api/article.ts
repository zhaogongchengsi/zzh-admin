import { Post, confirmStatus, Get } from "@/utils/service";
import type { article, state, articleList, ArticleType } from "@/types/response";
import type {
  articleStorageType,
  CosTempKeyRequest,
  article_req,
  tags,
} from "@/types/request";
import { putObject } from "./cos";
export async function getArticleList(
  limit?: number,
  offset?: number
): Promise<articleList> {
  return new Promise<articleList>((resolve, reject) => {
    Get(`article/get_article_list?limit=${limit}&offset=${offset}`)
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
  key: Array<string>
): Promise<boolean> {
  return new Promise<boolean>((resolve, reject) => {
    if (art.articleStorageType === articleStorageType.oos) {
      const cosOpt: CosTempKeyRequest = {
        Region: "ap-nanjing",
        Bucket: "bloghtml-1301735126",
        action: "",
      };
      const path = key.concat(art.articleName).join("/");
      putObject(cosOpt, art.articleContext, path, (err, data) => {
        if (err) {
          reject(err);
        } else {
          if (data.statusCode === 200) {
            art.articleUrl = data.Location;
            art.articleContext = "";
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

export async function getTagList(limit: number, offset: number): Promise<any> {
  return new Promise<any>((resolve, reject) => {
    Get(`article/tags/get_tags?limit=${limit}&offset=${offset}`)
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

export async function createTag(tag: tags): Promise<boolean> {
  return new Promise<boolean>((resolve, reject) => {
    Post("article/tags/create_tag", tag)
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

export async function getArticleType(): Promise<Array<ArticleType>> {
  return new Promise<Array<ArticleType>>((resolve, reject) => {
    Get(`article/types/get_types`)
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

export async function createArticleType (atype: ArticleType) : Promise<ArticleType> {
  return new Promise<ArticleType>((resolve, reject) => {
    Post('article/types/create_type', atype)
    .then(res => {
       const state = confirmStatus(<state>res.data.state);
        if (state) {
          resolve(res.data.data);
        } else {
          reject(res.data.state);
        }
    })
    .catch(reject)
  })
}
