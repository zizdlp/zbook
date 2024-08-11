import { Link } from "@/navigation";
export default function NotificationLink({
  redirect_url,
  children,
  onClickFunc,
}: {
  redirect_url: string;
  children: React.ReactNode;
  onClickFunc: any;
}) {
  if (redirect_url != "") {
    return (
      <Link
        href={redirect_url}
        onClick={onClickFunc}
        className="rounded-md md:rounded-xl dark:shadow-lg my-2 md:my-3
        bg-white dark:bg-slate-800 hover:dark:bg-sky-900 hover:bg-sky-500 hover:text-white border-[0.05rem] dark:border-0 flex items-center justify-between"
      >
        {children}
      </Link>
    );
  } else {
    return (
      <div
        onClick={onClickFunc}
        className="rounded-md md:rounded-xl dark:shadow-lg my-2 md:my-3
        bg-white dark:bg-slate-800 hover:dark:bg-sky-900 hover:bg-sky-500 hover:text-white border-[0.05rem] dark:border-0 flex items-center justify-between"
      >
        {children}
      </div>
    );
  }
}
