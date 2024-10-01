"use client";
import { useContext } from "react";
import { AiOutlineSearch } from "react-icons/ai";
import { SearchDialogContext } from "../SearchDialogProvider";
import { SearchType } from "@/utils/const_value";

export default function GlobalSearchButton() {
  const { setSearchDialogOpen, setSearchType } =
    useContext(SearchDialogContext);
  return (
    <AiOutlineSearch
      onClick={() => {
        setSearchType(SearchType.DOCUMENT);
        setSearchDialogOpen(true);
      }}
      aria-label="search"
      className="block w-6 h-6  hover:text-sky-600 dark:hover:text-sky-400 cursor-pointer"
    />
  );
}
