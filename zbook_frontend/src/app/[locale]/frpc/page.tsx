import FrpcSideBar from "./FrpcSideBar";

export default function frpc() {
  return (
    <div className="relative antialiased text-gray-500 dark:text-gray-400">
      <span className="fixed inset-0 bg-background-light dark:bg-background-dark"></span>
      <span className="fixed inset-0"></span>
      <div id="navbar" className="z-30 fixed top-0 w-full">
        <div className="max-w-8xl mx-auto">
          <div
            id="navbar-transition"
            className="backdrop-blur flex-none transition-colors duration-500 supports-backdrop-blur:bg-background-light/60 dark:bg-transparent"
          >
            <div className="relative lg:border-b lg:border-gray-500/5 dark:border-gray-50/[0.06]">
              <div className="flex items-center lg:px-12 border-b border-gray-500/10 lg:border-0 dark:border-gray-300/10 h-16 px-4">
                <div className="h-full relative flex-1 flex items-center gap-x-4">
                  <div className="flex-1 flex items-center gap-x-4">
                    <a href="/">
                      <img
                        className="w-auto h-7 relative object-contain block dark:hidden"
                        src="https://mintlify.s3-us-west-1.amazonaws.com/frpc/logo/light.svg"
                        alt="light logo"
                      />
                      <img
                        className="w-auto h-7 relative object-contain hidden dark:block"
                        src="https://mintlify.s3-us-west-1.amazonaws.com/frpc/logo/dark.svg"
                        alt="dark logo"
                      />
                    </a>
                  </div>
                  <div className="hidden lg:block mx-px relative flex-1 bg-white dark:bg-gray-900 pointer-events-auto rounded-lg">
                    <button
                      type="button"
                      className="w-full flex items-center text-sm leading-6 rounded-lg py-1.5 pl-2.5 pr-3 shadow-sm text-gray-400 dark:text-white/50 bg-background-light dark:bg-background-dark dark:brightness-[1.1] dark:ring-1 dark:hover:brightness-[1.25] ring-1 ring-gray-400/20 hover:ring-gray-600/25 dark:ring-gray-600/30 dark:hover:ring-gray-500/30 focus:outline-primary"
                      id="search-bar-entry"
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="16"
                        height="16"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        className="lucide lucide-search ml-1 mr-3 flex-none text-gray-700 hover:text-gray-800 dark:text-gray-300 hover:dark:text-gray-200"
                      >
                        <circle cx="11" cy="11" r="8"></circle>
                        <path d="m21 21-4.3-4.3"></path>
                      </svg>
                      Search or ask...
                      <span className="ml-auto flex-none text-xs font-semibold">
                        ⌘K
                      </span>
                    </button>
                  </div>
                  <div className="flex-1 relative hidden lg:flex items-center ml-auto justify-end space-x-4">
                    <nav className="text-sm">
                      <ul className="flex space-x-6 items-center">
                        <li className="cursor-pointer block lg:hidden">
                          <a
                            href="https://github.com/loopholelabs/frpc-go"
                            target="_blank"
                            rel="noreferrer"
                          >
                            <div className="group flex items-center space-x-3">
                              <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="1024"
                                height="1024"
                                viewBox="0 0 1024 1024"
                                fill="currentColor"
                                className="h-5 w-5"
                              >
                                <path
                                  fill-rule="evenodd"
                                  clip-rule="evenodd"
                                  d="M8 0C3.58 0 0 3.58 0 8C0 11.54 2.29 14.53 5.47 15.59C5.87 15.66 6.02 15.42 6.02 15.21C6.02 15.02 6.01 14.39 6.01 13.72C4 14.09 3.48 13.23 3.32 12.78C3.23 12.55 2.84 11.84 2.5 11.65C2.22 11.5 1.82 11.13 2.49 11.12C3.12 11.11 3.57 11.7 3.72 11.94C4.44 13.15 5.59 12.81 6.05 12.6C6.12 12.08 6.33 11.73 6.56 11.53C4.78 11.33 2.92 10.64 2.92 7.58C2.92 6.71 3.23 5.99 3.74 5.43C3.66 5.23 3.38 4.41 3.82 3.31C3.82 3.31 4.49 3.1 6.02 4.13C6.66 3.95 7.34 3.86 8.02 3.86C8.7 3.86 9.38 3.95 10.02 4.13C11.55 3.09 12.22 3.31 12.22 3.31C12.66 4.41 12.38 5.23 12.3 5.43C12.81 5.99 13.12 6.7 13.12 7.58C13.12 10.65 11.25 11.33 9.47 11.53C9.76 11.78 10.01 12.26 10.01 13.01C10.01 14.08 10 14.94 10 15.21C10 15.42 10.15 15.67 10.55 15.59C13.71 14.53 16 11.53 16 8C16 3.58 12.42 0 8 0Z"
                                  transform="scale(64)"
                                ></path>
                              </svg>
                              <div className="font-normal">
                                <div className="text-sm font-medium text-gray-700 group-hover:text-gray-900 dark:text-gray-300 dark:group-hover:text-gray-200">
                                  loopholelabsfrpc-go
                                </div>
                                <div className="text-xs flex items-center space-x-2 text-gray-600 dark:text-gray-400 group-hover:text-gray-700 dark:group-hover:text-gray-300">
                                  <span className="flex items-center space-x-1">
                                    <svg className="h-3 w-3 bg-gray-600 dark:bg-gray-400 group-hover:bg-gray-700 dark:group-hover:bg-gray-300"></svg>
                                    <span>439</span>
                                  </span>
                                  <span className="flex items-center space-x-1">
                                    <svg className="h-3 w-3 bg-gray-600 dark:bg-gray-400 group-hover:bg-gray-700 dark:group-hover:bg-gray-300"></svg>
                                    <span>17</span>
                                  </span>
                                </div>
                              </div>
                            </div>
                          </a>
                        </li>
                        <li className="cursor-pointer hidden lg:flex">
                          <a
                            href="https://github.com/loopholelabs/frpc-go"
                            target="_blank"
                            rel="noreferrer"
                          >
                            <div className="group flex items-center space-x-3">
                              <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="1024"
                                height="1024"
                                viewBox="0 0 1024 1024"
                                fill="currentColor"
                                className="h-5 w-5"
                              >
                                <path
                                  fill-rule="evenodd"
                                  clip-rule="evenodd"
                                  d="M8 0C3.58 0 0 3.58 0 8C0 11.54 2.29 14.53 5.47 15.59C5.87 15.66 6.02 15.42 6.02 15.21C6.02 15.02 6.01 14.39 6.01 13.72C4 14.09 3.48 13.23 3.32 12.78C3.23 12.55 2.84 11.84 2.5 11.65C2.22 11.5 1.82 11.13 2.49 11.12C3.12 11.11 3.57 11.7 3.72 11.94C4.44 13.15 5.59 12.81 6.05 12.6C6.12 12.08 6.33 11.73 6.56 11.53C4.78 11.33 2.92 10.64 2.92 7.58C2.92 6.71 3.23 5.99 3.74 5.43C3.66 5.23 3.38 4.41 3.82 3.31C3.82 3.31 4.49 3.1 6.02 4.13C6.66 3.95 7.34 3.86 8.02 3.86C8.7 3.86 9.38 3.95 10.02 4.13C11.55 3.09 12.22 3.31 12.22 3.31C12.66 4.41 12.38 5.23 12.3 5.43C12.81 5.99 13.12 6.7 13.12 7.58C13.12 10.65 11.25 11.33 9.47 11.53C9.76 11.78 10.01 12.26 10.01 13.01C10.01 14.08 10 14.94 10 15.21C10 15.42 10.15 15.67 10.55 15.59C13.71 14.53 16 11.53 16 8C16 3.58 12.42 0 8 0Z"
                                  transform="scale(64)"
                                ></path>
                              </svg>
                              <div className="font-normal">
                                <div className="text-sm font-medium text-gray-700 group-hover:text-gray-900 dark:text-gray-300 dark:group-hover:text-gray-200">
                                  loopholelabsfrpc-go
                                </div>
                                <div className="text-xs flex items-center space-x-2 text-gray-600 dark:text-gray-400 group-hover:text-gray-700 dark:group-hover:text-gray-300">
                                  <span className="flex items-center space-x-1">
                                    <svg className="h-3 w-3 bg-gray-600 dark:bg-gray-400 group-hover:bg-gray-700 dark:group-hover:bg-gray-300"></svg>
                                    <span>439</span>
                                  </span>
                                  <span className="flex items-center space-x-1">
                                    <svg className="h-3 w-3 bg-gray-600 dark:bg-gray-400 group-hover:bg-gray-700 dark:group-hover:bg-gray-300"></svg>
                                    <span>17</span>
                                  </span>
                                </div>
                              </div>
                            </div>
                          </a>
                        </li>
                      </ul>
                    </nav>
                    <div className="flex items-center">
                      <button className="group p-2 flex items-center justify-center">
                        <svg
                          width="16"
                          height="16"
                          viewBox="0 0 16 16"
                          fill="none"
                          stroke="currentColor"
                          xmlns="http://www.w3.org/2000/svg"
                          className="h-4 w-4 block text-gray-400 dark:hidden group-hover:text-gray-600"
                        >
                          <g clip-path="url(#clip0_2880_7340)">
                            <path
                              d="M8 1.11133V2.00022"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                            <path
                              d="M12.8711 3.12891L12.2427 3.75735"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                            <path
                              d="M14.8889 8H14"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                            <path
                              d="M12.8711 12.8711L12.2427 12.2427"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                            <path
                              d="M8 14.8889V14"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                            <path
                              d="M3.12891 12.8711L3.75735 12.2427"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                            <path
                              d="M1.11133 8H2.00022"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                            <path
                              d="M3.12891 3.12891L3.75735 3.75735"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                            <path
                              d="M8.00043 11.7782C10.0868 11.7782 11.7782 10.0868 11.7782 8.00043C11.7782 5.91402 10.0868 4.22266 8.00043 4.22266C5.91402 4.22266 4.22266 5.91402 4.22266 8.00043C4.22266 10.0868 5.91402 11.7782 8.00043 11.7782Z"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                          </g>
                          <defs>
                            <clipPath id="clip0_2880_7340">
                              <rect width="16" height="16" fill="white"></rect>
                            </clipPath>
                          </defs>
                        </svg>
                        <svg
                          width="16"
                          height="16"
                          viewBox="0 0 16 16"
                          fill="none"
                          stroke="currentColor"
                          xmlns="http://www.w3.org/2000/svg"
                          className="h-4 w-4 hidden dark:block text-gray-500 dark:group-hover:text-gray-300"
                        >
                          <g clip-path="url(#clip0_2880_7355)">
                            <path
                              d="M11.5556 10.4445C8.48717 10.4445 6.00005 7.95743 6.00005 4.88899C6.00005 3.68721 6.38494 2.57877 7.03294 1.66943C4.04272 2.22766 1.77783 4.84721 1.77783 8.0001C1.77783 11.5592 4.66317 14.4445 8.22228 14.4445C11.2196 14.4445 13.7316 12.3948 14.4525 9.62321C13.6081 10.1414 12.6187 10.4445 11.5556 10.4445Z"
                              stroke-width="1.5"
                              stroke-linecap="round"
                              stroke-linejoin="round"
                            ></path>
                          </g>
                          <defs>
                            <clipPath id="clip0_2880_7355">
                              <rect width="16" height="16" fill="white"></rect>
                            </clipPath>
                          </defs>
                        </svg>
                      </button>
                    </div>
                  </div>
                  <div className="flex lg:hidden items-center">
                    <button
                      type="button"
                      className="ml-auto text-gray-500 w-8 h-8 -my-1 items-center justify-center hover:text-gray-600 dark:text-gray-400 dark:hover:text-gray-300"
                      id="search-bar-entry-mobile"
                    >
                      <span className="sr-only">Search</span>
                      <svg className="h-4 w-4 bg-gray-500 dark:bg-gray-400 hover:bg-gray-600 dark:hover:bg-gray-300"></svg>
                    </button>
                    <button className="h-7 w-5 flex items-center justify-end">
                      <svg className="h-4 w-4 bg-gray-500 dark:bg-gray-400 hover:bg-gray-600 dark:hover:bg-gray-300"></svg>
                    </button>
                  </div>
                </div>
              </div>
              <div className="flex items-center h-14 py-4 px-5 border-b border-gray-500/10 lg:hidden dark:border-gray-50/[0.06]">
                <button
                  type="button"
                  className="text-gray-500 hover:text-gray-600 dark:text-gray-400 dark:hover:text-gray-300"
                >
                  <span className="sr-only">Navigation</span>
                  <svg
                    className="h-4"
                    fill="currentColor"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 448 512"
                  >
                    <path d="M0 96C0 78.3 14.3 64 32 64H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H32C14.3 128 0 113.7 0 96zM0 256c0-17.7 14.3-32 32-32H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H32c-17.7 0-32-14.3-32-32zM448 416c0 17.7-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32H416c17.7 0 32 14.3 32 32z"></path>
                  </svg>
                </button>
                <div className="ml-4 flex text-sm leading-6 whitespace-nowrap min-w-0 space-x-3">
                  <div className="flex items-center space-x-3">
                    <span>Welcome</span>
                    <svg
                      width="3"
                      height="24"
                      viewBox="0 -9 3 24"
                      className="h-5 rotate-0 overflow-visible fill-gray-400"
                    >
                      <path
                        d="M0 0L3 3L0 6"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="1.5"
                        stroke-linecap="round"
                      ></path>
                    </svg>
                  </div>
                  <div className="font-semibold text-gray-900 truncate dark:text-gray-200">
                    fRPC Documentation
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="px-4 mx-auto max-w-[92rem] lg:px-8">
        <FrpcSideBar />
        <div className="max-w-6xl" id="content-container">
          <div className="flex flex-row items-stretch gap-12 pt-[9.5rem] lg:pt-[6.5rem]">
            <div
              className="relative grow overflow-hidden mx-auto px-1 lg:-ml-12 lg:pl-[23.7rem]"
              id="content-area"
            >
              <header id="header" className="relative">
                <div className="mt-0.5 space-y-2.5">
                  <div className="eyebrow h-5 text-primary dark:text-primary-light text-sm font-semibold">
                    Welcome
                  </div>
                  <div className="flex items-center">
                    <h1 className="inline-block text-2xl sm:text-3xl font-extrabold text-gray-900 tracking-tight dark:text-gray-200">
                      fRPC Documentation
                    </h1>
                  </div>
                </div>
              </header>
              <div className="flex flex-col gap-8">
                <div className="flex flex-col gap-6 xl:hidden"></div>
              </div>
              <div className="relative mt-8 prose prose-gray dark:prose-invert">
                <p>
                  <button data-state="closed">
                    <span className="underline decoration-dotted decoration-2 underline-offset-4 decoration-gray-400 dark:decoration-gray-500">
                      fRPC
                    </span>
                  </button>{" "}
                  is a <strong>proto3-compatible</strong> RPC Framework that’s
                  designed from the ground up to be{" "}
                  <strong>
                    lightweight, extensible, and extremely performant
                  </strong>
                  . On average fRPC outperforms other RPC frameworks{" "}
                  <a href="/performance/grpc-benchmarks">
                    by 2-4x in an apples-to-apples comparison
                  </a>
                  , and is easily able to handle more than 2 million RPCs/second
                  on a single server.
                </p>
                <span aria-owns="rmiz-modal-3deaf47cc47c" data-rmiz="">
                  <span data-rmiz-content="found">
                    <img
                      className="w-full"
                      src="https://mintlify.s3-us-west-1.amazonaws.com/frpc/images/intro.svg"
                      alt="Welcome to fRPC"
                    />
                  </span>
                  <span data-rmiz-ghost="">
                    <button
                      aria-label="Expand image: Welcome to fRPC"
                      data-rmiz-btn-zoom=""
                      type="button"
                    >
                      <svg
                        aria-hidden="true"
                        data-rmiz-btn-zoom-icon="true"
                        fill="currentColor"
                        focusable="false"
                        viewBox="0 0 16 16"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path d="M 9 1 L 9 2 L 12.292969 2 L 2 12.292969 L 2 9 L 1 9 L 1 14 L 6 14 L 6 13 L 2.707031 13 L 13 2.707031 L 13 6 L 14 6 L 14 1 Z"></path>
                      </svg>
                    </button>
                  </span>
                </span>
                <div className="not-prose grid gap-x-4 sm:grid-cols-2">
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    href="/getting-started/overview"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Getting Started
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            Quickly get up and running with fRPC by following
                            our getting started guide.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    href="/getting-started/concepts"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Concepts
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            Take a look at some unique fRPC concepts and how it
                            differs from other frameworks.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    href="/getting-started/concepts"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Roadmap
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            Check out our planned technical roadmap to see how
                            we’ll be improving fRPC in the future.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    href="/reference/overview"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Technical Docs
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            Take a look at our technical docs to dig into the
                            details of fRPC and how you can use it.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                </div>
                <div className="not-prose grid gap-x-4 sm:grid-cols-2">
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    target="_blank"
                    rel="noreferrer"
                    href="https://github.com/loopholelabs/frpc-go"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Contributing
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            The footer of each page contains an “Edit on Github”
                            link. Please feel free to contribute by making a
                            pull request!
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    target="_blank"
                    rel="noreferrer"
                    href="https://loopholelabs.io/discord"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Join Our Discord
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            The <strong>#Frisbee</strong> Channel in our Discord
                            server is a great place to get help with all things
                            Frisbee and fRPC.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                </div>
              </div>
              <div className="leading-6 mt-14">
                <div className="mb-12 px-0.5 flex items-center text-sm font-semibold text-gray-700 dark:text-gray-200">
                  <a
                    className="flex items-center ml-auto space-x-3 group"
                    href="/getting-started/overview"
                  >
                    <span className="group-hover:text-gray-900 dark:group-hover:text-white">
                      Overview
                    </span>
                    <svg
                      viewBox="0 0 3 6"
                      className="rotate-180 h-1.5 stroke-gray-400 overflow-visible group-hover:stroke-gray-600 dark:group-hover:stroke-gray-300"
                    >
                      <path
                        d="M3 0L0 3L3 6"
                        fill="none"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      ></path>
                    </svg>
                  </a>
                </div>
                <footer className="justify-between pt-10 border-t border-gray-100 sm:flex dark:border-gray-800/50 pb-28">
                  <div className="flex mb-6 space-x-6 sm:mb-0">
                    <a href="https://loopholelabs.io/discord" target="_blank">
                      <span className="sr-only">discord</span>
                      <svg className="w-5 h-5 bg-gray-400 dark:bg-gray-500 hover:bg-gray-500 dark:hover:bg-gray-400"></svg>
                    </a>
                    <a
                      href="https://github.com/loopholelabs/frpc-go"
                      target="_blank"
                    >
                      <span className="sr-only">github</span>
                      <svg className="w-5 h-5 bg-gray-400 dark:bg-gray-500 hover:bg-gray-500 dark:hover:bg-gray-400"></svg>
                    </a>
                  </div>
                  <div className="sm:flex">
                    <a
                      href="https://mintlify.com?utm_campaign=poweredBy&amp;utm_medium=docs&amp;utm_source=frpc.io"
                      target="_blank"
                      rel="noreferrer"
                      className="text-sm text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-300"
                    >
                      Powered by Mintlify
                    </a>
                  </div>
                </footer>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
