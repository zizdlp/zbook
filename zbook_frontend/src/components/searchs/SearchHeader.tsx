import React, { useRef, useEffect } from "react";
import { MdCancel, MdPersonSearch } from "react-icons/md";
import { useTranslations } from "next-intl";
import SwitchType from "./SwitchType";
import { AiOutlineFileSearch } from "react-icons/ai";
import { SearchType } from "@/utils/const_value";

type SearchHeaderProps = {
  showDialog: boolean;
  setShowDialog: React.Dispatch<React.SetStateAction<boolean>>;
  query: string;
  setquery: React.Dispatch<React.SetStateAction<string>>;
  searchType: SearchType;
};

export default function SearchHeader(props: SearchHeaderProps) {
  const searchQueryRef = useRef<HTMLInputElement>(null);
  const t = useTranslations("Searchs");

  useEffect(() => {
    if (searchQueryRef.current) {
      searchQueryRef.current.value = props.query;
    }
  }, [props.query]);

  const clearquery = async (
    event: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    event.preventDefault();
    if (searchQueryRef.current) {
      searchQueryRef.current.value = "";
    }
    props.setquery("");
    props.setShowDialog(!props.showDialog);
  };

  const submitContact = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault(); // 阻止redirect to search query
    if (searchQueryRef.current) {
      props.setquery(searchQueryRef.current.value);
    } else {
      props.setquery("");
    }
  };

  let title = "";
  if (
    props.searchType == SearchType.USER ||
    props.searchType == SearchType.VISI_USER
  ) {
    title = t("SearchUser");
  } else {
    title = t("SearchMarkdown");
  }

  return (
    <header className="px-4 py-4 relative flex text-slate-500 flex-row items-center border-b-[0.05rem] border-slate-300 dark:border-slate-700/75">
      <form
        onChange={submitContact}
        onSubmit={submitContact}
        className="flex flex-grow flex-shrink items-center"
      >
        {(props.searchType === SearchType.DOCUMENT ||
          props.searchType === SearchType.USER) && <SwitchType />}

        {(props.searchType === SearchType.REPO_DOCUMENT ||
          props.searchType === SearchType.USER_DOCUMENT) && (
          <AiOutlineFileSearch className="h-7 w-7 " />
        )}

        {props.searchType === SearchType.VISI_USER && (
          <MdPersonSearch className="h-7 w-7" />
        )}

        <input
          ref={searchQueryRef}
          className="px-4 py-2 flex flex-grow flex-shrink focus:ring-0 focus-within:text-slate-900 focus-within:dark:text-slate-300 bg-transparent 
                placeholder-slate-400 text-gray-900 dark:text-slate-300 appearance-none w-full  border-0  focus:outline-none"
          id="search_query"
          type="text"
          name="search_query"
          placeholder={title}
        />
      </form>
      <button onClick={clearquery}>
        <MdCancel className="w-8 h-8 px-1.5 py-1 " />
      </button>
    </header>
  );
}
