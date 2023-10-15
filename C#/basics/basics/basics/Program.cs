using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CS_Program { 
    
    class ListSearching
    {
        static void Main(string[] args)
        {
            int[] list = { 1, 2, 3, 4, 5, 6, 7, 8 };
            int targetNum = 8;
            int targetIndex = BinarySearch(list, targetNum);

            Console.WriteLine("The index of list item is: " + targetIndex);
        }

        private static int BinarySearch(int[] list, int targetNum)
        {
            int biggerIndex = list.Length - 1;
            int lowestIndex = 0;

            for (int i = 0; i < list.Length; i++)
            {
               int midIndex = (biggerIndex + lowestIndex) / 2;

               if (list[midIndex] == targetNum)
               {
                   return midIndex;
               }

               if (list[midIndex] > targetNum)
               {
                   biggerIndex -= midIndex; 
               }

               if (list[midIndex] < targetNum)
               {
                   lowestIndex += midIndex;
               }
            }
            
            return -1;
        }
    }
}