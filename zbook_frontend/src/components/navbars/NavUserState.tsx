import { signOut } from "@/auth";
import { Link } from "@/navigation";
import { RiLogoutBoxLine, RiLoginBoxLine } from "react-icons/ri";
import { auth } from "@/auth";
import { AiOutlineSearch, AiOutlineUser } from "react-icons/ai";
import GlobalSearchButton from "@/providers/dialogs/GlobalSearchButton";
export default async function UserState() {
  const session = await auth();
  if (session && session.access_token) {
    return (
      <>
        <form
          className="flex"
          action={async () => {
            "use server";
            await signOut();
          }}
        >
          <button type="submit">
            <RiLogoutBoxLine className="block w-6 h-6  hover:text-sky-600 dark:hover:text-sky-400 cursor-pointer" />
          </button>
        </form>
        <GlobalSearchButton />
        <Link href={`/workspace/${session.username}`}>
          <AiOutlineUser className="block w-6 h-6  hover:text-sky-600 dark:hover:text-sky-400 cursor-pointer" />
        </Link>
      </>
    );
  } else {
    return (
      <>
        <Link href={`/auth/login`}>
          <RiLoginBoxLine className="block w-6 h-6  hover:text-sky-600 dark:hover:text-sky-400 cursor-pointer" />
        </Link>
      </>
    );
  }
}
