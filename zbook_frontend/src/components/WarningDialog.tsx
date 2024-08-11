import DialogComponent from "@/components/DialogComponent";
export default function WarngingDialog({
  title,
  showDialog,
  setShowDialog,
  cancelFunc,
  submitFunc,
  cancelTitle,
  submitTitle,
}: {
  title: string;
  showDialog: boolean;
  setShowDialog: React.Dispatch<React.SetStateAction<boolean>>;
  cancelFunc: any;
  submitFunc: any;
  cancelTitle: string;
  submitTitle: string;
}) {
  return (
    <DialogComponent showDialog={showDialog} setShowDialog={setShowDialog}>
      <header className="justify-center px-4 py-4 overflow-auto relative flex  text-slate-700 flex-row items-center border-b border-slate-300/75 dark:border-slate-800/75">
        <div className="flex lg:block md:overflow-hidden grow  justify-center overflow-auto  items-center ">
          <div className="sm:overflow-hidden sm:rounded-md">
            <div className="grid grid-cols-6 gap-4 gap-x-8 p-2">
              <div className="col-span-6 justify-center flex-center">
                <label
                  htmlFor="post_title"
                  className="block text-center pt-2 pb-6 text-2xl font-bold text-gray-700 dark:text-slate-50"
                >
                  {title}
                </label>
              </div>
            </div>
            <div className="grid grid-cols-2 gap-x-4 sm:gap-x-6 lg:gap-x-4 xl:gap-x-6 p-4 sm:px-6 sm:py-5 lg:p-4 xl:px-6 xl:py-5">
              <button
                className="text-base font-medium rounded-lg bg-white border-[0.05rem] dark:border-0 border-slate-300/50 text-slate-900 py-3 text-center cursor-pointer dark:bg-gray-700/75 dark:text-slate-200 dark:highlight-white/10"
                onClick={cancelFunc}
              >
                {cancelTitle}
              </button>
              <button
                className="text-base font-medium rounded-lg bg-sky-500 dark:bg-sky-800 text-white py-3 text-center cursor-pointer dark:highlight-white/20"
                onClick={submitFunc}
              >
                {submitTitle}
              </button>
            </div>
          </div>
        </div>
      </header>
    </DialogComponent>
  );
}
