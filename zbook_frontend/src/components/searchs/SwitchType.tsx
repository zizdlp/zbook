import { useContext } from "react";
import { MdPersonSearch } from "react-icons/md";
import { AiOutlineFileSearch } from "react-icons/ai";
import { SearchDialogContext } from "@/providers/SearchDialogProvider";
import { SearchType } from "@/utils/const_value";

export default function SwitchType() {
  const { setSearchType, searchType } = useContext(SearchDialogContext);

  return (
    <div className="relative inline-flex h-8 w-20 px-2 items-center rounded-full bg-gray-200/75 dark:bg-gray-900/75 cursor-pointer">
      <div
        onClick={() => {
          setSearchType(SearchType.DOCUMENT);
        }}
        className={`absolute left-1 h-7 w-7 rounded-full flex items-center justify-center transition-colors ${
          searchType === SearchType.DOCUMENT && "bg-white dark:bg-slate-700/75"
        }`}
      >
        <AiOutlineFileSearch className="h-6 w-6 p-1 " />
      </div>
      <div
        onClick={() => {
          setSearchType(SearchType.USER);
        }}
        className={`absolute right-1 h-7 w-7 rounded-full flex items-center justify-center transition-colors ${
          searchType === SearchType.USER && "bg-white dark:bg-slate-700/75"
        }`}
      >
        <MdPersonSearch className="h-6 w-6 p-1" />
      </div>
    </div>
  );
}
