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
  code: number
  message: string
  success: boolean
}

export interface VerificationCode {
  id: string
  b64s: string
} 