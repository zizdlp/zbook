import { Link } from "@/navigation";
import { ListUserInfo } from "@/types/model";
import AvatarImageClient from "../AvatarImageClient";
interface ProfileProps {
  ListUserInfo: ListUserInfo;
}

export default function SearchProfileComponent(props: ProfileProps) {
  return (
    <div
      className="rounded-md md:rounded-xl dark:shadow-lg my-2 md:my-3 py-2
        bg-white dark:bg-slate-800 hover:dark:bg-sky-900 hover:bg-sky-500 hover:text-white border-[0.05rem] dark:border-0 flex items-center justify-between"
    >
      <Link
        href={"/workspace/" + props.ListUserInfo.username}
        className="flex items-center justify-start px-2"
      >
        <AvatarImageClient
          username={props.ListUserInfo.username}
          className="w-12 h-12 mx-4 rounded-full shadow-lg"
        />
        <div className="flex flex-col py-2">
          <strong className=" text-sm font-medium  flex space-x-2 items-center justify-start">
            <p>{props.ListUserInfo.username}</p>
          </strong>
          <span className=" text-sm font-medium ">
            {props.ListUserInfo.email}
          </span>
        </div>
      </Link>
    </div>
  );
}
