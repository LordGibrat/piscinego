#!/bin/bash
export INTERVIEW_NUM=$(sed -n '179p' mystery/streets/Buckingham_Place | cut -d '#' -f2)
echo "$INTERVIEW_NUM"
cat "mystery/interviews/interview-${INTERVIEW_NUM}"
echo "$MAIN_SUSPECT"