import Image from "next/image";
export default function UserSideBarProfile({
  avatar,
  email,
  username,
  motto,
  bio,
}: {
  avatar: string;
  email: string;
  username: string;
  motto: string;
  bio: string;
}) {
  return (
    <>
      <div className="flex flex-col items-center justify-center p-4">
        {avatar ? (
          <Image
            src={`data:image/png;base64,${avatar}`}
            width={80}
            height={80}
            alt="Picture of the user"
            className={`flex-none w-24 h-24 rounded-full  object-cover`}
          />
        ) : (
          <div className="flex-none w-24 h-24 rounded-full  object-cover bg-slate-300 dark:bg-slate-700" />
        )}

        <p className="text-lg font-semibold">{username} </p>
        <p className="mt-0.5">{email}</p>
      </div>

      <blockquote className="text-slate-700 dark:text-slate-300 text-sm">
        <p
          className="overflow-x-auto  scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-h-[6px]
              scrollbar-thumb-slate-200 scrollbar-track-slate-100
              dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
        >
          <span className="font-bold pr-2 text-sm">{bio}</span>
          {motto}
        </p>
      </blockquote>
    </>
  );
}
