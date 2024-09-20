import { Dialog, Transition } from "@headlessui/react";
import React, { Fragment } from "react";
type DialogCompentProps = {
  children: React.ReactNode;
  showDialog: boolean;
  setShowDialog: React.Dispatch<React.SetStateAction<boolean>>;
};
export default function DialogComponent(props: DialogCompentProps) {
  function closeModal() {
    props.setShowDialog(false);
  }
  return (
    <Transition appear show={props.showDialog} as={Fragment}>
      <Dialog as="div" className="relative z-50" onClose={closeModal}>
        <div className="fixed inset-0 bg-transparent backdrop-blur-sm backdrop-brightness-75 transition-opacity" />
        <Transition.Child
          as={Fragment}
          enter="ease-out duration-300"
          enterFrom="opacity-0 scale-75"
          enterTo="opacity-100 scale-100"
          leave="ease-in duration-300"
          leaveFrom="opacity-100 scale-100"
          leaveTo="opacity-0 scale-50"
        >
          <div
            className="fixed inset-0 overflow-y-auto
          scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-w-[6px]
        scrollbar-thumb-slate-200 scrollbar-track-slate-100
        dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]
          dark:text-slate-300 md:p-20 flex flex-col justify-start pt-4 md:justify-center lg:p-28"
          >
            <div className="mx-4 md:mx-[max(0px,calc(50%-25rem))] ">
              <Dialog.Panel className="flex overflow-hidden flex-col  border rounded-md md:rounded-lg bg-white/85 dark:bg-slate-800/85 border-white/85 dark:border-slate-800/85">
                {props.children}
              </Dialog.Panel>
            </div>
          </div>
        </Transition.Child>
      </Dialog>
    </Transition>
  );
}
