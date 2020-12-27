# kdWatchDog

[![Build Status](https://github.com/petershen0307/kdWatchDog/workflows/Go/badge.svg?branch=master)](https://github.com/petershen0307/kdWatchDog/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/petershen0307/kdWatchDog)](https://goreportcard.com/report/github.com/petershen0307/kdWatchDog)

## KD 公式
[KD公式](https://www.ezchart.com.tw/inds.php?IND=KD)
[範例](http://yhhuang1966.blogspot.com/2015/02/kd.html)
>KD 隨機指標
>說明	KD市場常使用的一套技術分析工具。其適用範圍以中短期投資的技術分析為最佳。隨機指標的理論認為：當股市處於牛市時，收盤價往往接近當日最高價； 反之在熊市時，收盤價比較接近當日最低價，該指數的目的即在反映出近期收盤價在該段日子中價格區間的相對位置。
>計算公式
>它是由%K(快速平均值)、%D(慢速平均值)兩條線所組成，假設從n天週期計算出隨機指標時，首先須找出最近n天當中曾經出現過的最高價、最低價與第n天的收盤價，然後利用這三個數字來計算第n天的未成熟隨機值(RSV)
>         第n天收盤價-最近n天內最低價
>RSV ＝────────────────×100
>      最近n天內最高價-最近n天內最低價
>計算出RSV之後，再來計算K值與D值。
>當日K值(%K)= 2/3 前一日 K值 + 1/3 RSV
>當日D值(%D)= 2/3 前一日 D值＋ 1/3 當日K值
>**若無前一日的K值與D值，可以分別用50來代入計算**，經過長期的平滑的結果，起算基期雖然不同，但會趨於一致，差異很小。
>使用方法
>如果行情是一個明顯的漲勢，會帶動K線與D線向上升。如漲勢開始遲緩，則會反應到K值與D值，使得K值>跌破D值，此時中短期跌勢確立。
>當K值大於D值，顯示目前是向上漲升的趨勢，因此在圖形上K線向上突破D線時，即為買進訊號。
>當D值大於K值，顯示目前是向下跌落，因此在圖形上K 線向下跌破D線，此即為賣出訊號。
>上述K線與D線的交叉，須在80以上，20以下(一說70、30；視市場投機程度而彈性擴大範圍)，訊號才正確。
>當K值大於80，D值大於70時，表示當日收盤價處於偏高之價格區域，即為超買狀態；當K值小於20，D值>小於30時，表示當日收盤價處於偏低之價格區域，即為超賣狀態。
>當D值跌至15以下時，意味市場為嚴重之超賣，其為買入訊號；當D值超過85以上時，意味市場為嚴重之超買，其為賣出訊號。
>價格創新高或新低，而KD未有此現象，此為背離現象，亦即為可能反轉的重要前兆。
