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
        <p className="overflow-x-auto">
          <span className="font-bold pr-2 text-sm">{bio}</span>
          {motto}
        </p>
      </blockquote>
    </>
  );
}
