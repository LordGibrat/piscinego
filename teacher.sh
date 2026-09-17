INTERVIEW_NUM=$(sed -n '179p' streets/Buckingham_Place | cut -d '#' -f2)
echo "$INTERVIEW_NUM"
cat "interviews/interview-${INTERVIEW_NUM}"
echo "$MAIN_SUSPECT"
