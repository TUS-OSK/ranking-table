import { useLayoutEffect, useState } from 'react'
import { Crown } from 'lucide-react'

const useWindowSize = (): number[] => {
  const [size, setSize] = useState([0, 0]);
  useLayoutEffect(() => {
    const updateSize = (): void => {
      setSize([window.innerWidth, window.innerHeight]);
    };

    window.addEventListener('resize', updateSize);
    updateSize();

    return () => window.removeEventListener('resize', updateSize);
  }, []);
  return size;
};

export default function Leaderboard() {
  const [width, height] = useWindowSize();

  const topUsers = [
    { name: "田中太郎", score: 9850 },
    { name: "佐藤花子", score: 9720 },
    { name: "鈴木一郎", score: 9580 },
    { name: "高橋美咲", score: 9450 },
    { name: "伊藤健太", score: 9320 },
  ];

  const getCrownColor = (index: number) => {
    switch (index) {
      case 0: return 'text-yellow-400';
      case 1: return 'text-gray-400';
      case 2: return 'text-yellow-600';
      default: return '';
    }
  };

  const calculateSize = (baseSize: number, index: number): number => {
    const scaleFactor = Math.min(width, height) / 1200;
    return baseSize * scaleFactor * (3 - index * 0.1);
  };

  return (
    <div className="flex items-center justify-center w-screen h-screen bg-gradient-to-br from-blue-500 to-purple-600 overflow-hidden">
      <div className="w-full h-full flex flex-col bg-white">
        <h1 className="text-3xl font-bold text-center text-gray-800 py-2 bg-gray-100 flex-shrink-0"
          style={{ height: `${height * 0.1}px`, fontSize: `${calculateSize(36, 0)}px` }}>
          トップスコア
        </h1>
        <ul className="flex-grow flex flex-col justify-between p-2 overflow-hidden" role="list">
          {topUsers.map((user, index) => (
            <li
              key={index}
              className={`flex items-center justify-between rounded-lg transition-transform hover:scale-105
                ${index === 0 ? 'bg-yellow-100 shadow-lg' :
                  index === 1 ? 'bg-gray-200' :
                    index === 2 ? 'bg-gray-100' :
                      'bg-white'}`}
              style={{ height: `${height * (0.16 - index * 0.02)}px` }}
              role="listitem"
            >
              <div className="flex items-center h-full flex-1 px-2">
                <div className="relative mr-2" style={{ height: '80%', aspectRatio: '1 / 1' }}>
                  <span
                    className="absolute inset-0 font-bold text-white bg-blue-500 rounded-full flex items-center justify-center"
                    style={{ fontSize: `${calculateSize(24, index)}px` }}
                    aria-hidden="true"
                  >
                    {index + 1}
                  </span>
                  {index < 3 && (
                    <Crown
                      className={`absolute -top-10% -left-10% ${getCrownColor(index)}`}
                      style={{ width: '50%', height: '50%' }}
                      aria-hidden="true"
                    />
                  )}
                </div>
                <span className="font-bold text-gray-800 flex-1 px-2" style={{ fontSize: `${calculateSize(22, index)}px` }}>
                  {user.name}
                </span>
                <span
                  className={`font-bold ${index === 0 ? 'text-blue-600' :
                      index === 1 ? 'text-blue-500' :
                        index === 2 ? 'text-blue-400' :
                          'text-blue-300'
                    }`}
                  style={{ fontSize: `${calculateSize(24, index)}px` }}
                  aria-label={`スコア: ${user.score.toLocaleString()}`}
                >
                  {user.score.toLocaleString()}
                </span>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}