
export interface user {
  useradmin: string
  password: string
  captcha:string
  captchaId: string
}

export interface state {
  code: number
  message: string
  success: boolean
}

export interface article_tags {
	tag:string
	details:string
}

export enum articleStorageType {
  oos = "oos",
  databases = "databases",
  service = "service"
}

export interface article_req {
  fileName: string
  articleName: string
  articleTitle: string
  articleUrl: string
  articleStorageType: string
  articleAuthor: string
  articleType: string
  articleContext: string
  articleTags: Array<article_tags>
  article_desc: string
}

export interface CosTempKeyRequest {
  Region: string
  Bucket: string
  action: string
}

export interface tags {
  tag:string
	details: string
}
