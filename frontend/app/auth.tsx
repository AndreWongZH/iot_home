"use client"

export const Auth = (
    { isLogin, setLogin, setUsername, setPassword, tryLogin, register }:
    { isLogin: Boolean, setLogin: Function, setUsername: Function, setPassword: Function, tryLogin: Function, register: Function }
) => {

  return (
    <div className="flex flex-col justify-start items-center">
      <h1 className="font-bold text-2xl mb-20 pt-24">{isLogin ? "Welcome to IOT home" : "Register an account"}</h1>
      <div className="max-w-xs min-h-min mx-auto bg-highlight mt-16 p-7 rounded-lg">
        <div className="flex flex-col justify-center items-center gap-y-2.5">
          <input
            type="text"
            name="username"
            id="username"
            placeholder="Username"
            onChange={(e) => { setUsername(e.target.value) }}
            className="rounded-md border border-[#e0e0e0] bg-white py-3 px-3 w-full text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
          />

          <input
            type="password"
            name="password"
            id="password"
            placeholder="Password"
            onChange={(e) => { setPassword(e.target.value) }}
            className="rounded-md border border-[#e0e0e0] bg-white py-3 w-full px-3 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
          />

          <button
            className="w-full hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-3 text-center text-base font-semibold text-white outline-none"
            onClick={(e) => { isLogin ? tryLogin(e) : register(e) }}
          >{isLogin ? "Login" : "Register"}</button>
          <h1 className="font-bold text-2xl">OR</h1>
          <button
            className="w-full hover:shadow-lg rounded-md bg-gray-500 py-3 text-center text-base font-semibold text-white outline-none"
            onClick={() => { setLogin(!isLogin) }}
          >{isLogin ? "Register" : "Login"}</button>
        </div>
      </div>
    </div>
  )
}