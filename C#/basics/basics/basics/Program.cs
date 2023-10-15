using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CS_Program {
    public class ArraySorted
    {
        private List<int> list = new List<int>();

        public void populateList()
        {
            for (int i = 0; i < 10; i++)
            {
                list.Add(i);
            }
        }

        public List<int> getList() { return list; }

        public int[] toArray() { return list.ToArray(); }
    }

    class ListSearching
    {
        static void Main(string[] args)
        {
            ArraySorted listInstance = new ArraySorted();
            listInstance.populateList();

            int[] list = listInstance.getList().ToArray();
            int rand = new Random().Next(1, 10);

            int targetNum = rand;
            int targetIndex = BinarySearch(list, targetNum);

            Console.WriteLine("The index of list item is: " + targetIndex);
        }

        private static int BinarySearch(int[] list, int targetNum)
        {
            int biggerIndex = list.Length - 1;
            int lowestIndex = 0;

            while (lowestIndex <= biggerIndex)
            {
               int midIndex = (lowestIndex + biggerIndex) / 2;

                Console.WriteLine(midIndex);

               if (list[midIndex] == targetNum)
               {
                   return midIndex;
               }

               if (list[midIndex] > targetNum)
               {
                   biggerIndex = midIndex - 1; 
               }

               if (list[midIndex] < targetNum)
               {
                   lowestIndex = midIndex + 1;
               }
            }
            
            return -1;
        }
    }
}