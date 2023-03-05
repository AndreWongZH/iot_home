'use client'; // Error components must be Client components

import { useRouter } from 'next/navigation';
import { useEffect } from 'react';


export default function Error({
  error,
  reset,
}: {
  error: Error;
  reset: () => void;
  }) {
  const router = useRouter();
  useEffect(() => {
    // Log the error to an error reporting service
    console.log(error.message);
  }, [error]);

  return (
    <div className='w-full h-full flex flex-col items-center justify-center'>
      <h1 className='font-bold text-2xl'>Something went wrong!</h1>
      <h2 className='text-m mt-3'>{error.message}</h2>
      <button
        className="w-1/2 hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-10 text-center text-base font-semibold text-white outline-none"
        onClick={() => { router.replace(`/dashboard/`) }}
      >
        Reload
      </button>
    </div>
  );
}