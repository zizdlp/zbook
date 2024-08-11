import { FaGithub } from "react-icons/fa";
import { SiGitee } from "react-icons/si";
import { FaGitlab } from "react-icons/fa6";
import { FaGit } from "react-icons/fa6";

export default function GitHost({
  git_host,
  className,
}: {
  git_host: string;
  className: string;
}) {
  if (git_host.includes("github")) {
    return (
      <FaGithub className={`${className} text-gray-800 dark:text-gray-400`} />
    );
  } else if (git_host.includes("gitee")) {
    return (
      <SiGitee className={`${className} text-red-500 dark:text-gray-400`} />
    );
  } else if (git_host.includes("gitlab")) {
    return <FaGitlab className={`${className} `} />;
  } else {
    return <FaGit className={`${className} `} />;
  }
}
