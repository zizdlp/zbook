import { Link } from "@/navigation";
export default function RepoSideBarButton({
  onclick,
  className,
  url,
  title,
  children,
}: {
  onclick: any;
  className: string;
  url: string;
  title: string;
  children: React.ReactNode;
}) {
  if (url != "" && url != "#") {
    return (
      <Link
        onClick={onclick}
        href={url}
        className="group flex items-center lg:leading-6 mb-5 sm:mb-4  text-gray-600 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300 cursor-pointer"
      >
        <div
          className={`mr-4 rounded-md p-1 zinc-box group-hover:brightness-100 group-hover:ring-0 ring-1 ring-gray-950/5 dark:ring-gray-700/40 ${className}`}
        >
          {children}
        </div>
        {title}
      </Link>
    );
  } else {
    return (
      <div
        onClick={onclick}
        className="group flex items-center lg:leading-6 mb-5 sm:mb-4  text-gray-600 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300 cursor-pointer"
      >
        <div
          className={`mr-4 rounded-md p-1 zinc-box group-hover:brightness-100 group-hover:ring-0 ring-1 ring-gray-950/5 dark:ring-gray-700/40 ${className}`}
        >
          {children}
        </div>
        {title}
      </div>
    );
  }
}
