const emailRegex =
  /^(([^<>()$$$$\\.,;:\s@"]+(\.[^<>()$$$$\\.,;:\s@"]+)*)|(".+"))@(($$[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$$)|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

export { emailRegex };

// 定义枚举类型
export enum SearchType {
  DOCUMENT = 0,
  USER = 1,
  USER_DOCUMENT = 2,
  REPO_DOCUMENT = 3,
  VISI_USER = 4,
}
