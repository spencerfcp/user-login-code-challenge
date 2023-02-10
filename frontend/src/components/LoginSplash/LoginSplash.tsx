import Image from "next/image";
import { PropsWithChildren } from "react";

const LoginSplash = (props: PropsWithChildren) => {
  return (
    <div className=" bg-white shadow-lg border border-gray-100 flex items-center flex-col gap-2 h-[36rem]  p-5 rounded-md w-full flex-grow md:flex-grow-0 md:max-w-[30rem]">
      <div className="relative w-44 h-10 mt-5">
        <Image priority src="/images/scoir.svg" fill={true} alt="Scoir Logo" />
      </div>
      <span className="text-gray-400"> User Login Challenge</span>
      {props.children}
    </div>
  );
};

export default LoginSplash;
