// 用户信息
export interface adminUser {
  uuid: string;
  useradmin: string;
  userName: string;
  nickName: string;
  sideMode: string;
  avatarImg: string;
  baseColor: string;
  activeColor: string;
  authorityId: string;
  parentnodeid: string;
}
export interface userInfo {
  user: adminUser;
  token: string;
}

export interface state {
  code: number;
  message: string;
  success: boolean;
}

export interface VerificationCode {
  id: string;
  b64s: string;
}

export interface article {
  article_name: string;
  article_title: string;
  article_url: string;
  article_storage_type: string;
  article_tags: string[];
  article_author: string;
  article_context: string;
  article_file_name: string;
  article_type: string;
  likes: number;
  number_views: number;
  uuid: string;
  ID: number | string;
  article_desc: string;
}

export interface limit_offset {
  limit: number;
  offset: number;
}

export interface articleList {
  article_list: Array<article>;
  count: number;
  limit_offset: limit_offset;
}

export interface cosTemKey {
  Expiration: string;
  ExpiredTime: number;
  RequestId: string;
  StartTime: number;
  Credentials: Credentials;
}
export interface Credentials {
  TmpSecretId: string;
  TmpSecretKey: string;
  Token: string;
}
export interface cosPutObj {
  ETag: string;
  Location: string;
  RequestId: string;
}

export interface file {
  CreatedAt: string;
  ID: number;
  UpdatedAt: string;
  coverage_times: number;
  file_broad_type: string;
  file_ext: string;
  file_name: string;
  file_specific_type: string;
  is_hash: boolean;
  overwrite: boolean;
  sava_path: string;
}
