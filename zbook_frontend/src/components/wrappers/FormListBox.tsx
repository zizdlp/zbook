import { Fragment } from "react";
import { Listbox, Transition } from "@headlessui/react";
import { CheckIcon, ChevronUpDownIcon } from "@heroicons/react/20/solid";

interface ListBoxOption {
  label: string;
  value: string;
}

export default function FormListBox({
  options,
  formik,
  nameKey,
}: {
  options: ListBoxOption[];
  formik: any;
  nameKey: string;
}) {
  const selectedOption = options.find(
    (option) => option.value === formik.values[nameKey]
  ) || { label: "", value: "" };

  const handleChange = (value: string) => {
    formik.setFieldValue(nameKey, value);
  };

  return (
    <Listbox value={selectedOption.value} onChange={handleChange}>
      <div className="relative">
        <Listbox.Button
          className="p-3 rounded-md border border-slate-300 dark:border-slate-500  dark:text-slate-400 grow bg-white  dark:bg-slate-800 w-full text-sm
      placeholder:text-slate-400/75 dark:placeholder:text-slate-500/75 placeholder:text-sm  placeholder:font-base "
        >
          <span className="block truncate">{selectedOption.label}</span>
          <span className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
            <ChevronUpDownIcon
              className="h-5 w-5 text-gray-400"
              aria-hidden="true"
            />
          </span>
        </Listbox.Button>
        <Transition
          as={Fragment}
          leave="transition ease-in duration-100"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <Listbox.Options
            className="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-slate-800 py-1 text-base shadow-lg ring-1 ring-black/5 focus:outline-none sm:text-sm
                    overflow-y-scroll scrollbar-thin scrollbar-thumb-rounded-md scrollbar-track-rounded-md z-50 dark:border dark:border-slate-700"
          >
            {options.map((option, optionIdx) => (
              <Listbox.Option
                key={optionIdx}
                className={({ active }) =>
                  `relative cursor-default select-none py-2 pl-10 pr-4 ${
                    active
                      ? "bg-sky-400 dark:bg-sky-700 text-white"
                      : "text-gray-900 dark:text-gray-200"
                  }`
                }
                value={option.value}
              >
                {({ selected }) => (
                  <>
                    <span
                      className={`block truncate ${
                        selected ? "font-medium" : "font-normal"
                      }`}
                    >
                      {option.label}
                    </span>
                    {selected ? (
                      <span className="absolute inset-y-0 left-0 flex items-center pl-3 text-sky-600">
                        <CheckIcon className="h-5 w-5" aria-hidden="true" />
                      </span>
                    ) : null}
                  </>
                )}
              </Listbox.Option>
            ))}
          </Listbox.Options>
        </Transition>
      </div>
    </Listbox>
  );
}
