'use client'; // Error components must be Client components

export default function Error({
  error,
}: {
  error: string;
  }) {

  return (
    <div className='w-full h-full flex flex-col items-center justify-center'>
      <h1 className='font-bold text-3xl text-slate-500'>Something went wrong!</h1>
      <h2 className='text-xl mt-3 text-slate-500'>{error}</h2>
    </div>
  );
}