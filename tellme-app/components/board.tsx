'use client';

interface BoardProps {
  content: string;
}

export default function Board({ content }: BoardProps) {
  const formatContent = (text: string) => {
    return text.split('\n').map((line, index) => (
      <span key={index}>
        {line}
        <br />
      </span>
    ));
  };

  return (
    <div className="w-[360px] h-[540px] flex items-center justify-center">
      <a className="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700 h-full w-full flex items-center justify-center">
        <p className="font-normal text-gray-700 dark:text-gray-400 text-center">
          {formatContent(content)}
        </p>
      </a>
    </div>
  );
}
