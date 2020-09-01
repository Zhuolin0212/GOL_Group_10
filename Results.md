# Methods
Each language's buggy code has two bugs. One is a compilation bug to test how easy it is to follow the compiler's error message and the second is a logical bug to check how easy it is to grasp and understand the syntax of the language.

For any type of analysis to happen, we have to first collect data!

Participants in the experiment will give us the following points:

Languages the participant is familiar with.
How familiar the participant is with Go, Perl and Ruby.
Rank the languages based on how easy it was to debug among the three languages.
While the participants are debuging the code in each language we will gather the following data:

Debug time taken in each language.
Number of times the participant ran code for each language.
Based on the data collected we will compare each language and draw our conclusions!

# Materials  
### Survey
We used MonkeySurvey to make a questionnaire and counted the languages known by the participants. And ask participants that if they have used Go, Perl, and Ruby to code. 

### Code with bugs  
Our task sets two bugs for each language, one is compile bug, and another is logic bug. Each language has bugs designed in the same place. So after debugging one of the languages the participantmight spend less time finding the bug's location while debugging the other two languages. Therefore, the order in which the participants tested the three languages affected how long they ended up spending and how they ranked the language in terms of difficulty. 

### Online ide  
Partcipants were provided with an online ide link 'https://www.onlinegdb.com/' and required to share their screen during the experienments.

### Data collected  
By monitering the parcipants debugging process, we recorded how long it took to debug for each different language and evaluated if they passed the test. We also counted how many times participants excute the code in order to see if the warning messages gave enough cue for participants to find the bug right away.
At the end of the questionnaire, there is a question ask "how similar GO/Perl/Ruby looked as compared to the languages you have worked with". This is designed to consider whether it affects the debug speed of the participants.

# Observations

# Conclusions

# Threats to validity  

The above conclusions are subject to the validity of the experiments that we conducted. Some of the threats to validity that we found are:  
- Insufficient data points, since we could only conduct experiment with 9 different participants. Also, we think even 10 participants would not have been sufficient to draw reliable and trusty conclusions. Next time, we can increase the size of data points helping us draw reliable conclusion.
- Maladroit calculation of time taken to debug the codes as we calculated the time manually rather than automating it. So, we will automate the calculation of debug time.
- Ineffectual bugs in the 3 codes, since we had the same exact compilation and logical bugs in all the 3 codes. So next time, we can have different bugs in different codes keeping the difficulty level of bugs in each language same.
- Inconsistent level of complexity of the bugs in all the 3 languages. Participants found the logical bug of Ruby to be quite difficult because of the style of writing. Here, we can keep the level of complexity for all the languages to be the same next time.
