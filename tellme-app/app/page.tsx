'use client';
import { useState } from 'react';
import Board from '@/components/board';
import useFaro from '@/hooks/useFaro';

export default function Home() {
  const [joke, setJoke] = useState('');
  useFaro();

  const fetchJoke = async () => {
    const apiUrl = process.env.NEXT_PUBLIC_BACKEND_URL + '/joke';
    try {
      const response = await fetch(apiUrl);
      const data = await response.json();
      setJoke(`${data.setup} \n\n ${data.punchline}`);
    } catch (error) {
      console.error('Failed to fetch joke from:', apiUrl, error);
      setJoke('Failed to fetch joke.');
    }
  };

  return (
    <div className="overflow-hidden bg-white py-24 sm:py-32">
      <div className="mx-auto max-w-7xl px-6 lg:px-8">
        <div className="mx-auto grid max-w-2xl grid-cols-1 gap-x-8 gap-y-16 sm:gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-2">
          <div className="lg:pr-8 lg:pt-4">
            <button
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full ml-2 mr-2"
              onClick={fetchJoke} // Set up the onClick event handler
            >
              Tell me a joke
            </button>
            {/* Other buttons can be placed here */}
          </div>
          <Board content={joke} />{' '}
        </div>
      </div>
    </div>
  );
}
