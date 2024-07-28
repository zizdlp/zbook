export function isValidGitURL(gitURL: string): boolean {
  try {
    const u = new URL(gitURL);

    // 检查协议是否是 http、https、git、ssh
    if (
      u.protocol !== "http:" &&
      u.protocol !== "https:" &&
      u.protocol !== "git:" &&
      u.protocol !== "ssh:"
    ) {
      return false;
    }

    // 检查路径是否以 .git 结尾
    if (!u.pathname.endsWith(".git")) {
      return false;
    }

    return true;
  } catch (error) {
    return false;
  }
}

export function isValidateRepoName(repoName: string): Boolean {
  if (repoName.length < 2 || repoName.length > 64) {
    return false;
  }

  // Characters not allowed in URLs, typically include: '/', '?', ':', '@', '&', '=', '+', '$', ',', '#'
  const illegalChars = /[\/?:@&=+$,#~%]/;
  if (illegalChars.test(repoName)) {
    false;
  }

  return true;
}
