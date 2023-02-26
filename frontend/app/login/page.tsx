"use-client"

export default function Page() {
  return (
    <div className="flex flex-col justify-start items-center w-full h-full">
      <h1 className="font-bold text-2xl mb-20 pt-24">Welcome to IOT home</h1>
      <div className="max-w-xs h-64 mx-auto bg-highlight mt-16 p-7 rounded-lg">
        <form action="" method="POST">
          <div className="flex flex-col justify-center items-center gap-y-2.5">
            <input
              type="text"
              name="username"
              id="username"
              placeholder="Username"
              className="rounded-md border border-[#e0e0e0] bg-white py-3 px-3 w-full text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
            />

            <input
              type="password"
              name="password"
              id="password"
              placeholder="Password"
              className="rounded-md border border-[#e0e0e0] bg-white py-3 w-full px-3 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
            />

            <button className="w-full hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-3 text-center text-base font-semibold text-white outline-none">
              Login
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}