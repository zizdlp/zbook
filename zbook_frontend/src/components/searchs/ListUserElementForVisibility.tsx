import UpdateVisibleButton from "../wrappers/UpdateVisibilityButton";
import AvatarImageClient from "../AvatarImageClient";
export default function ListUserElementForVisibility({
  username,
  repo_username,
  repo_name,
  is_visible,
  email,
  is_following,
  repo_count,
  updated_at,
}: {
  username: string;
  repo_username: string;
  repo_name: string;
  is_visible: boolean;
  email: string;
  is_following: boolean;
  repo_count: number;
  updated_at: string;
}) {
  return (
    <div
      className="rounded-md md:rounded-xl dark:shadow-lg my-2 md:my-3 py-2
      bg-white dark:bg-slate-800 hover:dark:bg-sky-900 hover:bg-sky-500 hover:text-white border-[0.05rem] dark:border-0 flex items-center justify-between"
    >
      <div className="flex items-center justify-start px-2">
        <AvatarImageClient
          username={username}
          className="w-12 h-12 mx-4 rounded-full shadow-lg"
        />
        <div className="flex flex-col py-2">
          <strong className=" text-sm font-medium  flex space-x-2 items-center justify-start">
            <p>{username}</p>
          </strong>
          <span className=" text-sm font-medium ">{email}</span>
        </div>
      </div>
      <div className="pr-4">
        <UpdateVisibleButton
          username={username}
          repo_username={repo_username}
          repo_name={repo_name}
          is_visible={is_visible}
        />
      </div>
    </div>
  );
}
