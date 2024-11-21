import { useEffect, useLayoutEffect, useState } from 'react'
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

  const getCrownColor = (index: number) => {
    switch (index) {
      case 0: return 'text-yellow-400';
      case 1: return 'text-gray-400';
      case 2: return 'text-yellow-600';
      default: return '';
    }
  };

  const [topUsers, setTopUsers] = useState([
    { name: "田中太郎", score: 10003002 },
    { name: "佐藤花子", score: 9720 },
    { name: "鈴木一郎", score: 9580 },
    { name: "高橋美咲", score: 9450 },
    { name: "伊藤健太", score: 9320 },
  ]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://54.84.41.124:8080/');
        const data = await response.json();
        setTopUsers(data); // Update state with fetched data
      } catch (error) {
        console.error('Error fetching leaderboard data:', error);
      }
    };

    const intervalId = setInterval(fetchData, 1000);

    return () => clearInterval(intervalId);
  }, []);

  return (
    <div className="flex items-center justify-center w-screen h-screen bg-gradient-to-br from-blue-500 to-purple-600 overflow-hidden">
      <div className="w-full h-full flex flex-col bg-white">
        <div className="text-3xl font-bold text-center text-gray-800 py-2 mt-6 flex-shrink-0 flex items-center justify-center"
          style={{ height: `${height * 0.17}px`, fontSize: `${height * 0.09}px` }}>
          ⛳ Wikipedia Golf 本日のランキング ⛳
        </div>
        <ul className="flex-grow flex flex-col justify-between p-2 overflow-hidden" role="list">
          {topUsers.map((user, index) => (
            <li
              key={index}
              className={`flex items-center justify-between p-3`}
              style={{ height: `${height * (index === 0 ? 0.3 : index === 1 ? 0.2 : index === 2 ? 0.17 : 0.1)}px`, fontSize: `${height * (index === 0 ? 0.08 : index === 1 ? 0.06 : index === 2 ? 0.05 : 0.03)}px` }}
              role="listitem"
            >
              <div className={`h-full w-full flex items-center justify-between rounded-lg mx-4 my-2
          ${index === 0 ? 'bg-yellow-100 shadow-lg' :
                  index === 1 ? 'bg-gray-200 shadow-lg' :
                    index === 2 ? 'bg-orange-200 shadow-lg' :
                      'bg-white'}`}>
                <div className="flex items-center h-full flex-1 px-5">
                  <div className="relative mr-2" style={{ height: '80%', aspectRatio: '1 / 1' }}>
                    <span
                      className="absolute inset-0 font-bold text-white bg-blue-500 rounded-full flex items-center justify-center"
                      aria-hidden="true"
                    >
                      {index + 1}
                    </span>
                    {index < 3 && (
                      <Crown
                        className={`absolute -top-[10%] -left-[10%] ${getCrownColor(index)}`}
                        style={{ width: '55%', height: '55%' }}
                        aria-hidden="true"
                      />
                    )}
                  </div>
                  <span className="font-bold text-gray-800 flex-1 px-8" >
                    {user.name}
                  </span>
                  <span
                    className={`font-bold pr-10 text-blue-600`}
                    aria-label={`スコア: ${user.score.toLocaleString()}`}
                  >
                    {Math.floor(user.score / 10 ** 6) + '打  ' +
                      `${Math.floor((user.score % 10 ** 6) / 60).toString().padStart(2, '0')}:${(user.score % 60).toString().padStart(2, '0')}`}
                  </span>
                </div>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div >
  );
}