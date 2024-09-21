"use client";

import DialogComponent from "../../components/DialogComponent";
import React, { useContext, useState } from "react";
import SearchHeader from "../../components/searchs/SearchHeader";
import SearchFooter from "../../components/searchs/SearchFooter";
import { SearchDialogContext } from "@/providers/SearchDialogProvider";
import ListQueryElements from "@/components/searchs/ListQueryElements";

export default function SearchDialog() {
  const { searchDialogOpen, setSearchDialogOpen, searchType } =
    useContext(SearchDialogContext);
  const [query, setquery] = useState("");
  return (
    <DialogComponent
      showDialog={searchDialogOpen}
      setShowDialog={setSearchDialogOpen}
    >
      <SearchHeader
        showDialog={searchDialogOpen}
        setShowDialog={setSearchDialogOpen}
        query={query}
        setquery={setquery}
        searchType={searchType}
      />

      <div
        className="px-2 md:px-4 mx-4
             rounded-lg md:rounded-lg 
             ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2 overflow-x-clip
             overflow-y-auto  scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-w-[6px]
              scrollbar-thumb-slate-200 scrollbar-track-slate-100
              dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]
              md:h-96 h-[30rem] pb-2 md:pb-4"
      >
        <ListQueryElements
          key={query + searchType}
          searchType={searchType}
          query={query}
          setShowDialog={setSearchDialogOpen}
        />
      </div>
      <SearchFooter />
    </DialogComponent>
  );
}
